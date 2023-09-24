package main

import (
	"github.com/gofiber/fiber/v2"
	adapters "hexarch/adapters/driven"
	"hexarch/adapters/driving"
	"hexarch/core/services"
)

func main() {
	http := fiber.New()
	repo := adapters.NewUserRedisAdapter("localhost:6379")
	srv := services.NewUserService(repo)

	driving.SetupUserHTTPAdapter(srv, http.Group("/users"))
}
