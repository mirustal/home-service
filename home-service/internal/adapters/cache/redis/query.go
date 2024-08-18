package redis

import (
	"context"
	"fmt"
)

func (ra *RedisAdapter) Delete(key string) error {
	return ra.client.Del(context.Background(), key).Err()
}

func (ra *RedisAdapter) Set(key, value string) {
	ra.client.Set(context.Background(), key, []byte(value), ra.ttl)
}

func (ra *RedisAdapter) Get(key string) (string, string, error) {
	op := "adapters.cache.get"
	val, err := ra.client.Get(context.Background(), key).Result()

	if err != nil {
		ra.log.Error("%s: %v", op, err)
		return "", "", fmt.Errorf("%v", err)
	}

	return key, val, nil
}
