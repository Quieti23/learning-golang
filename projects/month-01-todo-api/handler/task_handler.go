package handler

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"

	"month01todoapi/service"
	"month01todoapi/store"
)

type TaskHandler struct {
	service *service.TaskService
}

type createTaskRequest struct {
	Title string `json:"title"`
}

type updateTaskRequest struct {
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

type errorResponse struct {
	Message string `json:"message"`
}

func NewTaskHandler(taskService *service.TaskService) *TaskHandler {
	return &TaskHandler{service: taskService}
}

func (h *TaskHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/tasks", h.handleTasks)
	mux.HandleFunc("/tasks/", h.handleTaskByID)
}

func (h *TaskHandler) handleTasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		writeJSON(w, http.StatusOK, h.service.List())
	case http.MethodPost:
		defer r.Body.Close()

		var request createTaskRequest
		if err := decodeJSONBody(r, &request); err != nil {
			writeErrorJSON(w, http.StatusBadRequest, err.Error())
			return
		}

		task, err := h.service.Create(request.Title)
		if err != nil {
			h.writeServiceError(w, err)
			return
		}

		writeJSON(w, http.StatusCreated, task)
	default:
		writeErrorJSON(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

func (h *TaskHandler) handleTaskByID(w http.ResponseWriter, r *http.Request) {
	id, err := parseTaskID(r.URL.Path)
	if err != nil {
		writeErrorJSON(w, http.StatusBadRequest, "invalid task id")
		return
	}

	switch r.Method {
	case http.MethodGet:
		h.handleGetTaskByID(w, id)
	case http.MethodPut:
		h.handleUpdateTask(w, r, id)
	case http.MethodDelete:
		h.handleDeleteTask(w, id)
	default:
		writeErrorJSON(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

func (h *TaskHandler) handleGetTaskByID(w http.ResponseWriter, id int) {
	task, err := h.service.GetByID(id)
	if err != nil {
		h.writeServiceError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, task)
}

func (h *TaskHandler) handleUpdateTask(w http.ResponseWriter, r *http.Request, id int) {
	defer r.Body.Close()

	var request updateTaskRequest
	if err := decodeJSONBody(r, &request); err != nil {
		writeErrorJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	task, err := h.service.Update(id, service.UpdateTaskInput{
		Title: request.Title,
		Done:  request.Done,
	})
	if err != nil {
		h.writeServiceError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, task)
}

func (h *TaskHandler) handleDeleteTask(w http.ResponseWriter, id int) {
	err := h.service.Delete(id)
	if err != nil {
		h.writeServiceError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *TaskHandler) writeServiceError(w http.ResponseWriter, err error) {
	switch err {
	case service.ErrTitleRequired:
		writeErrorJSON(w, http.StatusBadRequest, err.Error())
	case store.ErrTaskNotFound:
		writeErrorJSON(w, http.StatusNotFound, err.Error())
	default:
		writeErrorJSON(w, http.StatusInternalServerError, "internal server error")
	}
}

func writeErrorJSON(w http.ResponseWriter, statusCode int, message string) {
	writeJSON(w, statusCode, errorResponse{Message: message})
}

func writeJSON(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(data)
}

func decodeJSONBody(r *http.Request, dst any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(dst); err != nil {
		if errors.Is(err, io.EOF) {
			return errors.New("request body is required")
		}

		return errors.New("invalid JSON body")
	}

	if decoder.More() {
		return errors.New("request body must contain only one JSON object")
	}

	return nil
}

func parseTaskID(path string) (int, error) {
	idText := strings.TrimPrefix(path, "/tasks/")
	if idText == "" || strings.Contains(idText, "/") {
		return 0, strconv.ErrSyntax
	}

	return strconv.Atoi(idText)
}
