package handler

import (
	"gogpt/ui/components"
	"net/http"
)

type PageHandler struct {
}

func NewPageHandler() *PageHandler {
	return &PageHandler{}
}

func (h *PageHandler) HandleViewHomePage(w http.ResponseWriter, r *http.Request) {
	page := components.Home()
	page.Render(r.Context(), w)
}
