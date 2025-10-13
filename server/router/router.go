package router

import (
	"net/http"
	"time"

	"github.com/FrancoMusolino/go-todoapp/internal/api/handlers"
	"github.com/FrancoMusolino/go-todoapp/internal/domain/interfaces"
	"github.com/FrancoMusolino/go-todoapp/middlewares"
	"github.com/FrancoMusolino/go-todoapp/utils"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
)

const AUTH_ROUTER_MAX_REQUESTS_PER_MINUTE = 3

func SetUpRouter(authHandler *handlers.AuthHandler, todoHandler *handlers.TodoHandler, userRepo interfaces.IUserRepo) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer, middleware.RequestID, middleware.RealIP, middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://127.0.0.1:3000", "https://yappr.chat", "http://yappr.chat"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Route("/api/auth", func(u chi.Router) {
		u.Use(httprate.Limit(
			AUTH_ROUTER_MAX_REQUESTS_PER_MINUTE,
			time.Minute,
			httprate.WithKeyFuncs(httprate.KeyByIP, httprate.KeyByEndpoint),
			httprate.WithLimitHandler(func(w http.ResponseWriter, r *http.Request) {
				utils.WriteError(w, http.StatusTooManyRequests, "Demasiados intentos. Por favor, vuelva en unos minutos", nil)
				return
			}),
		))

		u.Post("/register", authHandler.Register)
		u.Post("/login", authHandler.Login)
		u.Post("/verify-user", authHandler.VerifyUser)
		u.Post("/resend-verification-email", authHandler.ResendVerificationEmail)
	})

	r.Route("/api/todos", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(middlewares.JWTAuth(userRepo))
			r.Get("/", todoHandler.GetUserTodos)
			r.Post("/", todoHandler.CreateTodo)
		})
	})

	r.Get("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	})

	return r
}
