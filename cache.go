package cacheredis

import (
	"time"

	"github.com/casbin/casbin/v2/persist/cache"
	"github.com/go-redis/redis"
)

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(client *redis.Client) *RedisCache {
	return &RedisCache{
		client: client,
	}
}

func (c *RedisCache) Set(key string, value []byte, expireTime time.Duration) error {
	c.client.Set(key, value, expireTime)
	return nil
}

// Get function returns cache.ErrNoSuchKey on any error, for avoid casbin downfall when redis isn't available.
func (c *RedisCache) Get(key string) ([]byte, error) {
	result := c.client.Get(key)

	if result.Err() == redis.Nil {
		// Key was not found in cache.
		return nil, cache.ErrNoSuchKey
	}

	if result.Err() != nil {
		// Some error in redis client.
		return nil, cache.ErrNoSuchKey
	}

	res, err := result.Bytes()

	if err != nil {
		return nil, cache.ErrNoSuchKey
	}

	return res, nil
}

func (c *RedisCache) Delete(key string) error {
	c.client.Del(key)
	return nil
}

func (c *RedisCache) Clear() error {
	c.client.FlushDB()
	return nil
}
