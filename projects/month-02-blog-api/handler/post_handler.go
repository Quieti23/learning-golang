package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
	"time"

	"month02blogapi/repository"
	"month02blogapi/service"
)

type PostHandler struct {
	service service.PostService
	logger  *slog.Logger
	timeout time.Duration
}

type createPostRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

type updatePostRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

type errorResponse struct {
	Message string `json:"message"`
}

func NewPostHandler(postService service.PostService, logger *slog.Logger, timeout time.Duration) *PostHandler {
	return &PostHandler{service: postService, logger: logger, timeout: timeout}
}

func (h *PostHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/posts", h.handlePosts)
	mux.HandleFunc("/posts/", h.handlePostByID)
}

func (h *PostHandler) handlePosts(w http.ResponseWriter, r *http.Request) {
	requestLogger := h.logger.With("method", r.Method, "path", r.URL.Path)

	switch r.Method {
	case http.MethodGet:
		ctx, cancel := h.withRequestTimeout(r.Context())
		defer cancel()

		posts, err := h.service.List(ctx)
		if err != nil {
			h.writeServiceError(requestLogger, w, err)
			return
		}
		requestLogger.Info("list posts", "count", len(posts))
		writeJSON(w, http.StatusOK, posts)
	case http.MethodPost:
		defer r.Body.Close()

		var request createPostRequest
		if err := decodeJSONBody(r, &request); err != nil {
			requestLogger.Warn("decode post request failed", "error", err)
			writeErrorJSON(w, http.StatusBadRequest, err.Error())
			return
		}

		ctx, cancel := h.withRequestTimeout(r.Context())
		defer cancel()

		post, err := h.service.Create(ctx, service.CreatePostInput{
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

func (h *PostHandler) handlePostByID(w http.ResponseWriter, r *http.Request) {
	requestLogger := h.logger.With("method", r.Method, "path", r.URL.Path)

	id, err := parsePostID(r.URL.Path)
	if err != nil {
		requestLogger.Warn("invalid post id", "error", err)
		writeErrorJSON(w, http.StatusBadRequest, "invalid post id")
		return
	}

	switch r.Method {
	case http.MethodGet:
		ctx, cancel := h.withRequestTimeout(r.Context())
		defer cancel()

		post, err := h.service.GetByID(ctx, id)
		if err != nil {
			h.writeServiceError(requestLogger, w, err)
			return
		}

		requestLogger.Info("get post by id", "post_id", id)
		writeJSON(w, http.StatusOK, post)
	case http.MethodPut:
		defer r.Body.Close()

		var request updatePostRequest
		if err := decodeJSONBody(r, &request); err != nil {
			requestLogger.Warn("decode update post request failed", "error", err)
			writeErrorJSON(w, http.StatusBadRequest, err.Error())
			return
		}

		ctx, cancel := h.withRequestTimeout(r.Context())
		defer cancel()

		post, err := h.service.Update(ctx, id, service.UpdatePostInput{
			Title:   request.Title,
			Content: request.Content,
			Author:  request.Author,
		})
		if err != nil {
			h.writeServiceError(requestLogger, w, err)
			return
		}

		requestLogger.Info("post updated", "post_id", id)
		writeJSON(w, http.StatusOK, post)
	case http.MethodDelete:
		ctx, cancel := h.withRequestTimeout(r.Context())
		defer cancel()

		if err := h.service.Delete(ctx, id); err != nil {
			h.writeServiceError(requestLogger, w, err)
			return
		}

		requestLogger.Info("post deleted", "post_id", id)
		w.WriteHeader(http.StatusNoContent)
	default:
		requestLogger.Warn("method not allowed")
		writeErrorJSON(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

func (h *PostHandler) writeServiceError(logger *slog.Logger, w http.ResponseWriter, err error) {
	logger.Warn("post request failed", "error", err)

	switch err {
	case service.ErrTitleRequired, service.ErrContentRequired, service.ErrAuthorRequired:
		writeErrorJSON(w, http.StatusBadRequest, err.Error())
	case repository.ErrPostNotFound:
		writeErrorJSON(w, http.StatusNotFound, err.Error())
	default:
		if errors.Is(err, context.DeadlineExceeded) {
			writeErrorJSON(w, http.StatusGatewayTimeout, "request timeout")
			return
		}

		if errors.Is(err, context.Canceled) {
			writeErrorJSON(w, http.StatusRequestTimeout, "request canceled")
			return
		}

		if errors.Is(err, repository.ErrPostNotFound) {
			writeErrorJSON(w, http.StatusNotFound, repository.ErrPostNotFound.Error())
			return
		}

		writeErrorJSON(w, http.StatusInternalServerError, "internal server error")
	}
}

func (h *PostHandler) withRequestTimeout(parent context.Context) (context.Context, context.CancelFunc) {
	if h.timeout <= 0 {
		return parent, func() {}
	}

	return context.WithTimeout(parent, h.timeout)
}

func parsePostID(path string) (int, error) {
	idText := strings.TrimPrefix(path, "/posts/")
	if idText == "" || strings.Contains(idText, "/") {
		return 0, strconv.ErrSyntax
	}

	return strconv.Atoi(idText)
}

func writeErrorJSON(w http.ResponseWriter, statusCode int, message string) {
	writeJSON(w, statusCode, errorResponse{Message: message})
}

func decodeJSONBody(r *http.Request, dst any) error {
	if r.Body == nil {
		return errors.New("request body is required")
	}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		return errors.New("read request body failed")
	}

	if len(bytes.TrimSpace(data)) == 0 {
		return errors.New("request body is required")
	}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(dst); err != nil {
		var syntaxError *json.SyntaxError
		var typeError *json.UnmarshalTypeError

		switch {
		case errors.As(err, &syntaxError):
			return errors.New("invalid JSON body")
		case errors.As(err, &typeError):
			return errors.New("invalid JSON field type")
		case errors.Is(err, io.EOF):
			return errors.New("request body is required")
		default:
			if strings.HasPrefix(err.Error(), "json: unknown field ") {
				return errors.New("request body contains unknown fields")
			}

			return errors.New("invalid JSON body")
		}
	}

	if decoder.More() {
		return errors.New("request body must contain only one JSON object")
	}

	var extra any
	if err := decoder.Decode(&extra); err != io.EOF {
		return errors.New("request body must contain only one JSON object")
	}

	return nil
}

func writeJSON(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(data)
}
