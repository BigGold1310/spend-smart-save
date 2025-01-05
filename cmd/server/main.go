package main

import (
	"github.com/BigGold1310/spend-smart-save/internal/handler/api"
	"github.com/BigGold1310/spend-smart-save/internal/handler/ui"
	"github.com/BigGold1310/spend-smart-save/internal/repository/local"
	"github.com/BigGold1310/spend-smart-save/internal/service"
	"github.com/uptrace/bunrouter"
	"log"
	"net/http"
)

func main() {
	// Initialize repositories
	userRepo := local.NewUserRepository()

	// Initialize services
	userService := service.NewUserService(userRepo)

	apiHandler := api.NewUserHandler(userService)
	uiHandler := ui.NewUserHandler(userRepo)

	router := bunrouter.New()

	// Register API routes
	router.GET("/api/users", apiHandler.GetUsers)
	router.GET("/api/users/:id", apiHandler.GetUserByID)
	router.POST("/api/users", apiHandler.CreateUser)
	router.PUT("/api/users/:id", apiHandler.UpdateUser)

	// Register UI routes
	router.GET("/users", uiHandler.ListUsers)
	router.GET("/users/create", uiHandler.ShowCreateForm)
	router.GET("/users/:id", uiHandler.ShowUser)

	// Start the server
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
