package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"month02blogapi/service"
)

type PostHandler struct {
	service service.PostService
	logger  *slog.Logger
}

type createPostRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

type errorResponse struct {
	Message string `json:"message"`
}

func NewPostHandler(postService service.PostService, logger *slog.Logger) *PostHandler {
	return &PostHandler{service: postService, logger: logger}
}

func (h *PostHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/posts", h.handlePosts)
}

func (h *PostHandler) handlePosts(w http.ResponseWriter, r *http.Request) {
	requestLogger := h.logger.With("method", r.Method, "path", r.URL.Path)

	switch r.Method {
	case http.MethodGet:
		posts := h.service.List()
		requestLogger.Info("list posts", "count", len(posts))
		writeJSON(w, http.StatusOK, posts)
	case http.MethodPost:
		defer r.Body.Close()

		var request createPostRequest
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			requestLogger.Warn("decode post request failed", "error", err)
			writeErrorJSON(w, http.StatusBadRequest, "invalid JSON body")
			return
		}

		post, err := h.service.Create(service.CreatePostInput{
			Title:   request.Title,
			Content: request.Content,
			Author:  request.Author,
		})
		if err != nil {
			requestLogger.Warn("create post failed", "error", err)
			switch err {
			case service.ErrTitleRequired, service.ErrContentRequired, service.ErrAuthorRequired:
				writeErrorJSON(w, http.StatusBadRequest, err.Error())
			default:
				writeErrorJSON(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		requestLogger.Info("post created", "post_id", post.ID, "author", post.Author)
		writeJSON(w, http.StatusCreated, post)
	default:
		requestLogger.Warn("method not allowed")
		writeErrorJSON(w, http.StatusMethodNotAllowed, "method not allowed")
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
