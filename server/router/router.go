package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func SetUpRouter() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)

	// r.Group("/api/auth", func(u chi.Router) {
	// 	u.Post("/login")
	// 	u.Post("/signup")
	// })

	return r
}
