package render

import (
	"context"
	"io"

	"github.com/labstack/echo/v4"
)

type Component interface {
	Render(ctx context.Context, w io.Writer) error
}

func Render(ctx echo.Context, comp Component) error {
	return comp.Render(ctx.Request().Context(), ctx.Response().Writer)
}
