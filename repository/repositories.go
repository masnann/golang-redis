package repository

import (
	"context"
	"encoding/json"
	"errors"
	"golangredis/models"

	"github.com/go-redis/redis/v8"
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

	err = r.rdb.Set(context.Background(), "app_anan", jsonData, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r RedisRepository) GetData(key string) (*models.RequestRedis, error) {
	val, err := r.rdb.Get(context.Background(), key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, errors.New("Key not found in Redis")
		}
		return nil, err
	}

	var data models.RequestRedis
	err = json.Unmarshal([]byte(val), &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}