package repository

import "golangredis/models"

type RedisRepositoryInterface interface {
	Insert(data models.RequestRedis) error
	GetData(key string) (*models.RequestRedis, error)
}