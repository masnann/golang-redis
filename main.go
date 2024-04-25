package main

import (
	"golangredis/db"
	"golangredis/repository"
	"golangredis/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	if err := db.RedisInit(); err != nil {
		panic(err)
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	redisRepo := repository.NewRedisRepository(db.RedisConnect())
	redisService := service.NewRedisService(redisRepo)

	// Routes
	e.POST("/api/v1/redis/insert", redisService.Insert)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
