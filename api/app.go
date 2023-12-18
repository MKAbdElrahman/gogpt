package api

import (
	"gogpt/api/handler"
	"gogpt/ui/html"
	"os"

	"github.com/labstack/echo/v4"

	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/gommon/log"
)

func NewApp() *echo.Echo {
	app := echo.New()

	if l, ok := app.Logger.(*log.Logger); ok {
		l.SetOutput(os.Stderr)
		l.SetLevel(log.INFO)
		l.SetHeader(`${level} ` +
			`${short_file}|${line}`)
	}

	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           "time=${time_custom}, method=${method}, uri=${uri}, status=${status}, latency=${latency_human}\n",
		CustomTimeFormat: "15:04:05",
	}))

	app.Use(middleware.Recover())

	pageHandler := handler.NewPageHandler()

	app.Static("/static", "ui/assets")

	errorHandler := handler.ErrorHandler{}
	app.HTTPErrorHandler = errorHandler.HandleErrors

	app.GET("/", pageHandler.HandleViewPage(html.Home()))

	return app
}
