package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	_ "github.com/swaggo/http-swagger/example/go-chi/docs"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server.

// @host localhost:8080
func main() {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	r.Use(middleware.Logger)
	r.Get("/", welcome)

	http.ListenAndServe(":3000", r)
}

// ShowAccount godoc
// @Summary welcome
// @Description welcome message
// @Accept  json
// @Produce  json
// @Success 200 {string} welcome
// @Router / [get]
func welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome"))
}
