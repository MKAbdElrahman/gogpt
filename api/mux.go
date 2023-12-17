package api

import (
	"gogpt/api/handler"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewMux() *chi.Mux {
	mux := chi.NewRouter()

	mux.Use(middleware.Logger)
	mux.NotFound(handler.HandleNotFound)
	mux.MethodNotAllowed(handler.HandleMethodNotAllowed)

	pageHandler := handler.NewPageHandler()

	mux.Get("/", pageHandler.HandleViewHomePage)

	return mux
}
