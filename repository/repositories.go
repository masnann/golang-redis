package repository

import (
	"encoding/json"
	"golangredis/models"

	"github.com/go-redis/redis"
)

type RedisRepository struct {
	rdb *redis.Client
}

func NewRedisRepository(rdb *redis.Client) RedisRepository {
	return RedisRepository{
		rdb: rdb,
	}

}

func (r RedisRepository) Insert(data models.RequestRedis) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	
	err = r.rdb.Set("app_anan", jsonData, 0).Err()
	if err != nil {
		return err

	}
	return nil

}
