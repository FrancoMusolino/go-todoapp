package router

import (
	"net/http"

	"github.com/FrancoMusolino/go-todoapp/internal/api/handlers"
	"github.com/FrancoMusolino/go-todoapp/middlewares"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

func SetUpRouter(authHandler *handlers.AuthHandler, todoHandler *handlers.TodoHandler) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://127.0.0.1:3000", "https://yappr.chat", "http://yappr.chat"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Route("/api/auth", func(u chi.Router) {
		u.Post("/register", authHandler.Register)
		u.Post("/login", authHandler.Login)
	})

	r.Route("/api/todo", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(middlewares.JWTAuth)
			r.Get("/todos", todoHandler.GetUserTodos)
			r.Post("/", todoHandler.CreateTodo)
		})
	})

	r.Get("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	})

	return r
}
