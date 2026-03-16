package handler

import (
	"encoding/json"
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
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			writeJSON(w, http.StatusBadRequest, errorResponse{Message: "invalid JSON body"})
			return
		}

		task, err := h.service.Create(request.Title)
		if err != nil {
			if err == service.ErrTitleRequired {
				writeJSON(w, http.StatusBadRequest, errorResponse{Message: err.Error()})
				return
			}

			writeJSON(w, http.StatusInternalServerError, errorResponse{Message: "internal server error"})
			return
		}

		writeJSON(w, http.StatusCreated, task)
	default:
		writeJSON(w, http.StatusMethodNotAllowed, errorResponse{Message: "method not allowed"})
	}
}

func (h *TaskHandler) handleTaskByID(w http.ResponseWriter, r *http.Request) {
	idText := strings.TrimPrefix(r.URL.Path, "/tasks/")
	if idText == "" || strings.Contains(idText, "/") {
		writeJSON(w, http.StatusBadRequest, errorResponse{Message: "invalid task id"})
		return
	}

	id, err := strconv.Atoi(idText)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, errorResponse{Message: "invalid task id"})
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
		writeJSON(w, http.StatusMethodNotAllowed, errorResponse{Message: "method not allowed"})
	}
}

func (h *TaskHandler) handleGetTaskByID(w http.ResponseWriter, id int) {
	task, err := h.service.GetByID(id)
	if err != nil {
		if err == store.ErrTaskNotFound {
			writeJSON(w, http.StatusNotFound, errorResponse{Message: err.Error()})
			return
		}

		writeJSON(w, http.StatusInternalServerError, errorResponse{Message: "internal server error"})
		return
	}

	writeJSON(w, http.StatusOK, task)
}

func (h *TaskHandler) handleUpdateTask(w http.ResponseWriter, r *http.Request, id int) {
	defer r.Body.Close()

	var request updateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeJSON(w, http.StatusBadRequest, errorResponse{Message: "invalid JSON body"})
		return
	}

	task, err := h.service.Update(id, service.UpdateTaskInput{
		Title: request.Title,
		Done:  request.Done,
	})
	if err != nil {
		switch err {
		case service.ErrTitleRequired:
			writeJSON(w, http.StatusBadRequest, errorResponse{Message: err.Error()})
		case store.ErrTaskNotFound:
			writeJSON(w, http.StatusNotFound, errorResponse{Message: err.Error()})
		default:
			writeJSON(w, http.StatusInternalServerError, errorResponse{Message: "internal server error"})
		}
		return
	}

	writeJSON(w, http.StatusOK, task)
}

func (h *TaskHandler) handleDeleteTask(w http.ResponseWriter, id int) {
	err := h.service.Delete(id)
	if err != nil {
		if err == store.ErrTaskNotFound {
			writeJSON(w, http.StatusNotFound, errorResponse{Message: err.Error()})
			return
		}

		writeJSON(w, http.StatusInternalServerError, errorResponse{Message: "internal server error"})
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func writeJSON(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(data)
}
