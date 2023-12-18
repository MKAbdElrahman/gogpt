package handler

import (
	"gogpt/api/handler/render"

	"github.com/labstack/echo/v4"
)

type StaticHandler struct {
}

func NewPageHandler() *StaticHandler {
	return &StaticHandler{}
}

func (h *StaticHandler) HandleViewPage(page render.Component) echo.HandlerFunc {
	return func(c echo.Context) error {
		return render.Render(c, page)
	}
}
