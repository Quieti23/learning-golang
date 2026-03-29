package handler

import "net/http"

const (
	ErrorCodeInvalidRequest   = "INVALID_REQUEST"
	ErrorCodeNotFound         = "NOT_FOUND"
	ErrorCodeMethodNotAllowed = "METHOD_NOT_ALLOWED"
	ErrorCodeRequestTimeout   = "REQUEST_TIMEOUT"
	ErrorCodeRequestCanceled  = "REQUEST_CANCELED"
	ErrorCodeInternalError    = "INTERNAL_ERROR"
)

type errorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func writeErrorJSON(w http.ResponseWriter, statusCode int, code, message string) {
	writeJSON(w, statusCode, errorResponse{
		Code:    code,
		Message: message,
	})
}

func writeInvalidRequestError(w http.ResponseWriter, message string) {
	writeErrorJSON(w, http.StatusBadRequest, ErrorCodeInvalidRequest, message)
}

func writeMethodNotAllowedError(w http.ResponseWriter) {
	writeErrorJSON(w, http.StatusMethodNotAllowed, ErrorCodeMethodNotAllowed, "method not allowed")
}

func writeInternalError(w http.ResponseWriter) {
	writeErrorJSON(w, http.StatusInternalServerError, ErrorCodeInternalError, "internal server error")
}
