package repository

import "golangredis/models"

type RedisRepositoryInterface interface {
	Insert(data models.RequestRedis) error
}