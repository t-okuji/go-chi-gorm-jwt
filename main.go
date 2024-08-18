package main

import (
	"net/http"

	_ "github.com/swaggo/http-swagger/example/go-chi/docs"
	"github.com/t-okuji/go-chi-gorm-jwt/controller"
	"github.com/t-okuji/go-chi-gorm-jwt/db"
	"github.com/t-okuji/go-chi-gorm-jwt/repository"
	"github.com/t-okuji/go-chi-gorm-jwt/router"
	"github.com/t-okuji/go-chi-gorm-jwt/usecase"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server.

// @host localhost:8080
func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)
	r := router.NewRouter(userController)

	http.ListenAndServe(":3000", r)
}
