package redis

import (
	"context"
	"encoding/json"

	"home-service/internal/models"
)

func (userCache *RedisAdapter) Delete(key string) error {
	return userCache.client.Del(context.Background(), key).Err()
}

func (ra *RedisAdapter) Set(key string, value *models.User) {
	op := "adapters.cache.set"

	json, err := json.Marshal(value)
	if err != nil {
	ra.log.Error("%s: %v", op, err)
	}	

	ra.client.Set(context.Background(), key, json, ra.ttl)
}

func (ra *RedisAdapter) Get(key string) *models.User {
	op := "adapters.cache.get"
	val, err := ra.client.Get(context.Background(), key).Result()
	
	if err != nil {
		ra.log.Error("%s: %v", op, err)
		}	

	user := models.User{}
	err = json.Unmarshal([]byte(val), &user)
	if err != nil {
		panic(err)
	}
	return &user
}
