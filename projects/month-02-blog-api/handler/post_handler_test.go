package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"month02blogapi/model"
	"month02blogapi/repository"
	"month02blogapi/service"
)

func TestPostHandlerCRUDFlow(t *testing.T) {
	h := newTestHandler()

	created := performJSONRequest[tCreatePostResponse](t, h, http.MethodPost, "/posts", `{"title":"first","content":"hello","author":"eson"}`)
	if created.statusCode != http.StatusCreated {
		t.Fatalf("expected status %d, got %d", http.StatusCreated, created.statusCode)
	}
	if created.body.ID != 1 {
		t.Fatalf("expected created id 1, got %d", created.body.ID)
	}

	got := performJSONRequest[model.Post](t, h, http.MethodGet, "/posts/1", "")
	if got.statusCode != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, got.statusCode)
	}
	if got.body.Title != "first" {
		t.Fatalf("expected title first, got %q", got.body.Title)
	}

	updated := performJSONRequest[model.Post](t, h, http.MethodPut, "/posts/1", `{"title":"updated","content":"new content","author":"eson"}`)
	if updated.statusCode != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, updated.statusCode)
	}
	if updated.body.Title != "updated" {
		t.Fatalf("expected updated title, got %q", updated.body.Title)
	}

	list := performJSONRequest[[]model.Post](t, h, http.MethodGet, "/posts", "")
	if list.statusCode != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, list.statusCode)
	}
	if len(list.body) != 1 {
		t.Fatalf("expected 1 post, got %d", len(list.body))
	}
	if list.body[0].Title != "updated" {
		t.Fatalf("expected list to contain updated post, got %q", list.body[0].Title)
	}

	deleted := performRequest(t, h, http.MethodDelete, "/posts/1", "")
	if deleted.Code != http.StatusNoContent {
		t.Fatalf("expected status %d, got %d", http.StatusNoContent, deleted.Code)
	}

	notFound := performJSONRequest[errorResponse](t, h, http.MethodGet, "/posts/1", "")
	if notFound.statusCode != http.StatusNotFound {
		t.Fatalf("expected status %d, got %d", http.StatusNotFound, notFound.statusCode)
	}
	if notFound.body.Message != repository.ErrPostNotFound.Error() {
		t.Fatalf("expected message %q, got %q", repository.ErrPostNotFound.Error(), notFound.body.Message)
	}
}

func TestPostHandlerRejectsUnknownFields(t *testing.T) {
	h := newTestHandler()

	response := performJSONRequest[errorResponse](t, h, http.MethodPost, "/posts", `{"title":"x","content":"y","author":"z","extra":"bad"}`)
	if response.statusCode != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, response.statusCode)
	}
	if response.body.Message != "request body contains unknown fields" {
		t.Fatalf("unexpected error message: %q", response.body.Message)
	}
}

func TestPostHandlerRejectsEmptyBody(t *testing.T) {
	h := newTestHandler()

	response := performJSONRequest[errorResponse](t, h, http.MethodPost, "/posts", "   ")
	if response.statusCode != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, response.statusCode)
	}
	if response.body.Message != "request body is required" {
		t.Fatalf("unexpected error message: %q", response.body.Message)
	}
}

func TestPostHandlerRejectsInvalidPostID(t *testing.T) {
	h := newTestHandler()

	response := performJSONRequest[errorResponse](t, h, http.MethodGet, "/posts/not-a-number", "")
	if response.statusCode != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, response.statusCode)
	}
	if response.body.Message != "invalid post id" {
		t.Fatalf("unexpected error message: %q", response.body.Message)
	}
}

type tCreatePostResponse struct {
	ID int `json:"id"`
}

type testJSONResponse[T any] struct {
	statusCode int
	body       T
}

func newTestHandler() *PostHandler {
	repo := repository.NewInMemoryPostRepository()
	svc := service.NewPostService(repo)
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	return NewPostHandler(svc, logger)
}

func performRequest(t *testing.T, h *PostHandler, method, path, body string) *httptest.ResponseRecorder {
	t.Helper()

	req := httptest.NewRequest(method, path, bytes.NewBufferString(body))
	if method == http.MethodPost || method == http.MethodPut {
		req.Header.Set("Content-Type", "application/json")
	}

	recorder := httptest.NewRecorder()
	mux := http.NewServeMux()
	h.RegisterRoutes(mux)
	mux.ServeHTTP(recorder, req)
	return recorder
}

func performJSONRequest[T any](t *testing.T, h *PostHandler, method, path, body string) testJSONResponse[T] {
	t.Helper()

	recorder := performRequest(t, h, method, path, body)

	var result T
	if recorder.Body.Len() > 0 {
		if err := json.Unmarshal(recorder.Body.Bytes(), &result); err != nil {
			t.Fatalf("unmarshal response failed: %v", err)
		}
	}

	return testJSONResponse[T]{
		statusCode: recorder.Code,
		body:       result,
	}
}