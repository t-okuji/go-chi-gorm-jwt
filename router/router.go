package router

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth/v5"
	"github.com/t-okuji/go-chi-gorm-jwt/controller"
)

func NewRouter(uc controller.IUserController) http.Handler {
	tokenAuth := jwtauth.New("HS256", []byte(os.Getenv("TOKEN_SECRET")), nil)
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	// r.Use(middleware.Logger)

	// Protected routes
	r.Group(func(r chi.Router) {
		// Seek, verify and validate JWT tokens
		r.Use(jwtauth.Verifier(tokenAuth))

		// Handle valid / invalid tokens.
		r.Use(jwtauth.Authenticator(tokenAuth))

		r.Get("/admin", func(w http.ResponseWriter, r *http.Request) {
			_, claims, _ := jwtauth.FromContext(r.Context())
			w.Write([]byte(fmt.Sprintf("user_id : %v", claims["user_id"])))
		})

		r.Post("/logout", uc.LogOut)
	})

	// Public routes
	r.Group(func(r chi.Router) {
		r.Post("/signup", uc.SignUp)
		r.Post("/login", uc.LogIn)
	})

	return r
}
