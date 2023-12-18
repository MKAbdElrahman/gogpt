package handler

import (
	"fmt"

	"gogpt/api/handler/render"
	"gogpt/ui/html"

	"github.com/labstack/echo/v4"
)

type ErrorHandler struct {
}

func NewErrorHandler() *ErrorHandler {
	return &ErrorHandler{}
}

func (h ErrorHandler) HandleErrors(err error, c echo.Context) {
	if he, ok := err.(*echo.HTTPError); ok {
		logMessage := fmt.Sprintf("HTTP %d Error: %v", he.Code, he.Message)
		if he.Internal != nil {
			logMessage += fmt.Sprintf(". Internal Error: %v", he.Internal)
		}

		c.Logger().Error(logMessage)

		page := html.Error(he.Code, he.Message.(string))

		err := render.Render(c, page)

		if err != nil {
			c.Logger().Error(err)
		}

	}
}
