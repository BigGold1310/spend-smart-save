package main

import (
	"fmt"
	api2 "github.com/BigGold1310/spend-smart-save/internal/handler/api"
	"github.com/BigGold1310/spend-smart-save/internal/repository/local"
	"github.com/uptrace/bunrouter"
	"log"
	"net/http"
)

func main() {
	router := bunrouter.New()
	setupUIRoutes(router)
	setupAPIRoutes(router)

	// Start the server
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func setupUIRoutes(router *bunrouter.Router) {
	ui := router.NewGroup("/ui")
	ui.GET("/", func(w http.ResponseWriter, req bunrouter.Request) error {
		// req embeds *http.Request and has all the same fields and methods
		fmt.Println(req.Method, req.Route(), req.Params().Map())
		return nil
	})
}

func setupAPIRoutes(router *bunrouter.Router) {
	api := router.NewGroup("/api")
	u := api2.UserHandler{U: local.NewUserRepository()}
	api.POST("/users", u.Create)
	api.GET("/users/:id", u.Get)
	api.GET("/users", u.Get)
	api.DELETE("/users/:id", u.Delete)
}
