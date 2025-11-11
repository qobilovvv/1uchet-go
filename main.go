package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/qobilovvv/1uchet/config"
	"github.com/qobilovvv/1uchet/handlers"
	"github.com/qobilovvv/1uchet/repositories"
	"github.com/qobilovvv/1uchet/services"
)

const (
	PORT = ":3000"
)

func main() {
	db := config.InitDB()
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Post("/users/create", userHandler.CreateUser)

	http.ListenAndServe(PORT, router)
}
