package service

import (
	"golangredis/models"
	"golangredis/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

type RedisService struct {
	repo repository.RedisRepositoryInterface
}

func NewRedisService(repo repository.RedisRepository) RedisService {
	return RedisService{
		repo: repo,
	}
}

func (s RedisService) Insert(ctx echo.Context) error {
	// Bind request data to models.RequestRedis
	var requestData models.RequestRedis
	if err := ctx.Bind(&requestData); err != nil {
		return err
	}

	// Insert data to Redis using repository
	err := s.repo.Insert(requestData)
	if err != nil {
		return err
	}

	// Return success response
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "Data inserted successfully to Redis",
	})
}

func (s RedisService) GetData(ctx echo.Context) error {
	key := ctx.Param("key")

	data, err := s.repo.GetData(key)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, data)
}