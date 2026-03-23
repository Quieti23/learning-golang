package middleware

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRequestLoggerLogsMethodPathStatusAndDuration(t *testing.T) {
	var logBuffer bytes.Buffer
	logger := slog.New(slog.NewJSONHandler(&logBuffer, nil))

	handler := RequestLogger(logger)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}))

	recorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	handler.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusNoContent {
		t.Fatalf("expected status %d, got %d", http.StatusNoContent, recorder.Code)
	}

	entry := parseSingleLogEntry(t, logBuffer.String())
	if entry["msg"] != "http request completed" {
		t.Fatalf("expected message %q, got %#v", "http request completed", entry["msg"])
	}
	if entry["method"] != http.MethodGet {
		t.Fatalf("expected method %q, got %#v", http.MethodGet, entry["method"])
	}
	if entry["path"] != "/ping" {
		t.Fatalf("expected path %q, got %#v", "/ping", entry["path"])
	}
	if entry["status"] != float64(http.StatusNoContent) {
		t.Fatalf("expected status %d, got %#v", http.StatusNoContent, entry["status"])
	}
	duration, ok := entry["duration"].(string)
	if !ok || duration == "" {
		t.Fatalf("expected non-empty duration string, got %#v", entry["duration"])
	}
}

func TestRequestLoggerDefaultsStatusToOK(t *testing.T) {
	var logBuffer bytes.Buffer
	logger := slog.New(slog.NewJSONHandler(&logBuffer, nil))

	handler := RequestLogger(logger)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("pong"))
	}))

	recorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	handler.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, recorder.Code)
	}

	entry := parseSingleLogEntry(t, logBuffer.String())
	if entry["status"] != float64(http.StatusOK) {
		t.Fatalf("expected status %d, got %#v", http.StatusOK, entry["status"])
	}
}

func parseSingleLogEntry(t *testing.T, logs string) map[string]any {
	t.Helper()

	lines := strings.Split(strings.TrimSpace(logs), "\n")
	if len(lines) != 1 {
		t.Fatalf("expected 1 log line, got %d", len(lines))
	}

	var entry map[string]any
	if err := json.Unmarshal([]byte(lines[0]), &entry); err != nil {
		t.Fatalf("unmarshal log entry failed: %v", err)
	}

	return entry
}
