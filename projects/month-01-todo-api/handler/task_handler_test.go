package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"month01todoapi/service"
	"month01todoapi/store"
)

func newTestHandler() *TaskHandler {
	return NewTaskHandler(service.NewTaskService(store.NewTaskStore()))
}

func createTaskForTest(t *testing.T, handler *TaskHandler, title string) {
	t.Helper()

	request := httptest.NewRequest(http.MethodPost, "/tasks", strings.NewReader(`{"title":"`+title+`"}`))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	handler.handleTasks(recorder, request)

	if recorder.Code != http.StatusCreated {
		t.Fatalf("got status %d, want %d", recorder.Code, http.StatusCreated)
	}
}

func TestHandleTasksCreateSuccess(t *testing.T) {
	handler := newTestHandler()
	request := httptest.NewRequest(http.MethodPost, "/tasks", strings.NewReader(`{"title":"learn handler test"}`))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	handler.handleTasks(recorder, request)

	if recorder.Code != http.StatusCreated {
		t.Fatalf("got status %d, want %d", recorder.Code, http.StatusCreated)
	}

	var response map[string]any
	if err := json.Unmarshal(recorder.Body.Bytes(), &response); err != nil {
		t.Fatalf("unmarshal response failed: %v", err)
	}

	if response["title"] != "learn handler test" {
		t.Fatalf("got title %v, want %q", response["title"], "learn handler test")
	}
}

func TestHandleTasksCreateInvalidBody(t *testing.T) {
	handler := newTestHandler()
	request := httptest.NewRequest(http.MethodPost, "/tasks", strings.NewReader(`{"title":123}`))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	handler.handleTasks(recorder, request)

	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("got status %d, want %d", recorder.Code, http.StatusBadRequest)
	}
}

func TestHandleTaskByID(t *testing.T) {
	handler := newTestHandler()
	createTaskForTest(t, handler, "learn get by id")

	tests := []struct {
		name       string
		method     string
		target     string
		wantStatus int
	}{
		{
			name:       "existing task returns ok",
			method:     http.MethodGet,
			target:     "/tasks/1",
			wantStatus: http.StatusOK,
		},
		{
			name:       "missing task returns not found",
			method:     http.MethodGet,
			target:     "/tasks/999",
			wantStatus: http.StatusNotFound,
		},
		{
			name:       "invalid task id returns bad request",
			method:     http.MethodGet,
			target:     "/tasks/abc",
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := httptest.NewRequest(test.method, test.target, nil)
			recorder := httptest.NewRecorder()

			handler.handleTaskByID(recorder, request)

			if recorder.Code != test.wantStatus {
				t.Fatalf("got status %d, want %d", recorder.Code, test.wantStatus)
			}
		})
	}
}

func TestHandleTaskByIDUpdate(t *testing.T) {
	handler := newTestHandler()
	createTaskForTest(t, handler, "learn update")

	tests := []struct {
		name       string
		target     string
		body       string
		wantStatus int
	}{
		{
			name:       "update existing task returns ok",
			target:     "/tasks/1",
			body:       `{"title":"learn update done","done":true}`,
			wantStatus: http.StatusOK,
		},
		{
			name:       "update missing task returns not found",
			target:     "/tasks/999",
			body:       `{"title":"missing task","done":true}`,
			wantStatus: http.StatusNotFound,
		},
		{
			name:       "update empty title returns bad request",
			target:     "/tasks/1",
			body:       `{"title":"   ","done":true}`,
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPut, test.target, strings.NewReader(test.body))
			request.Header.Set("Content-Type", "application/json")
			recorder := httptest.NewRecorder()

			handler.handleTaskByID(recorder, request)

			if recorder.Code != test.wantStatus {
				t.Fatalf("got status %d, want %d", recorder.Code, test.wantStatus)
			}
		})
	}
}

func TestHandleTaskByIDDelete(t *testing.T) {
	handler := newTestHandler()
	createTaskForTest(t, handler, "learn delete")

	deleteRequest := httptest.NewRequest(http.MethodDelete, "/tasks/1", nil)
	deleteRecorder := httptest.NewRecorder()
	handler.handleTaskByID(deleteRecorder, deleteRequest)

	if deleteRecorder.Code != http.StatusNoContent {
		t.Fatalf("got status %d, want %d", deleteRecorder.Code, http.StatusNoContent)
	}

	getDeletedRequest := httptest.NewRequest(http.MethodGet, "/tasks/1", nil)
	getDeletedRecorder := httptest.NewRecorder()
	handler.handleTaskByID(getDeletedRecorder, getDeletedRequest)

	if getDeletedRecorder.Code != http.StatusNotFound {
		t.Fatalf("got status %d, want %d", getDeletedRecorder.Code, http.StatusNotFound)
	}

	deleteMissingRequest := httptest.NewRequest(http.MethodDelete, "/tasks/999", nil)
	deleteMissingRecorder := httptest.NewRecorder()
	handler.handleTaskByID(deleteMissingRecorder, deleteMissingRequest)

	if deleteMissingRecorder.Code != http.StatusNotFound {
		t.Fatalf("got status %d, want %d", deleteMissingRecorder.Code, http.StatusNotFound)
	}
}
