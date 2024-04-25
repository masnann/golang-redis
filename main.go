package main

import (
	"golangredis/db"
	"golangredis/repository"
	"golangredis/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db.RedisInit()

	rdb := db.RedisConnect()
	redisRepo := repository.NewRedisRepository(rdb)
	redisService := service.NewRedisService(redisRepo)
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())



	// Routes
	e.POST("/api/v1/redis/insert", redisService.Insert)
	e.GET("/api/v1/redis/get/:key", redisService.GetData)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
