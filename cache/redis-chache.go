package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

type (
	Post struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}

	PostChace interface {
		Set(key string, value Post)error
		Get(key string) (*Post, error)
	}
)

type redisCache struct {
	client *redis.Client // Menyimpan klien Redis sebagai properti
	exp    time.Duration
}

func NewRedisChace(host string, db int, exp time.Duration) PostChace {

	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: "", // Anda dapat menambahkan kata sandi jika dibutuhkan
		DB:       db,
	})

	return &redisCache{
		client: client,
		exp:    exp,
	}
}

func (r *redisCache) Set(key string, value Post) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	ctx := context.Background()
	err = r.client.Set(ctx, key, data, r.exp).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *redisCache) Get(key string) (*Post, error) {
	ctx := context.Background()

	v, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return nil, err // Mengembalikan kesalahan daripada panic
	}
	post := Post{}
	err = json.Unmarshal([]byte(v), &post)
	if err != nil {
		return nil, err // Mengembalikan kesalahan daripada panic
	}
	return &post, nil
}
