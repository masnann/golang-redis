package db

import "github.com/go-redis/redis"

var rdb *redis.Client

func RedisInit() error {
	db := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := db.Ping().Result()
	if err != nil {
		return err
	}

	rdb = db
	return nil
}

func RedisConnect() *redis.Client {
	return rdb
}
