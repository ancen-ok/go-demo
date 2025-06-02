package core

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

var (
	Cache *CacheRedis
	ctx   = context.Background()
)

type CacheRedis struct {
	client *redis.Client
}

// SetKeyValue 设置 Key-Value，并设置过期时间
func (c *CacheRedis) SetKeyValue(key string, value any, expiration time.Duration) (string, error) {
	err := c.client.Set(ctx, key, value, expiration).Err()
	return key, err
}

// GetKey 获取 Key 对应的 Value
func (c *CacheRedis) GetKey(key string) (string, error) {
	return c.client.Get(ctx, key).Result()
}

// Delete 删除一个或多个 key
func (c *CacheRedis) Delete(keys ...string) error {
	return c.client.Del(ctx, keys...).Err()
}

// Exist 判断 key 是否存在
func (c *CacheRedis) Exist(key string) bool {
	val, err := c.client.Exists(ctx, key).Result()
	return err == nil && val == 1
}

// IsExpire 获取 key 的剩余存活时间（单位：秒）
func (c *CacheRedis) IsExpire(key string) float64 {
	val, err := c.client.TTL(ctx, key).Result()
	if err != nil {
		return -2
	}
	return val.Seconds()
}

// KeyExpired 为 key 续期
func (c *CacheRedis) KeyExpired(key string, duration time.Duration) (bool, error) {
	return c.client.Expire(ctx, key, duration).Result()
}

func InitRedis() {
	config := Config.Redis
	addr := fmt.Sprintf("%s:%d", config.Host, config.Port)
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: config.Password,
		DB:       config.Db,
	})

	if _, err := client.Ping(ctx).Result(); err != nil {
		Log.Error("redis 连接失败: %v", err.Error())
		panic(err.Error())
	}

	Log.Info("redis 连接成功")
	Cache = &CacheRedis{client: client}
}
