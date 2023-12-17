package handler

import "net/http"

type ErrorHandler struct {
}

func NewErrorHandler() *ErrorHandler {
	return &ErrorHandler{}
}

func (h *ErrorHandler) HandleBadRequest(w http.ResponseWriter, r *http.Request, message string) {

	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(message))
}

