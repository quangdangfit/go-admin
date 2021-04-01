package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/quangdangfit/gocommon/logger"

	"github.com/quangdangfit/go-admin/config"
)

// constants cache
const (
	RedisExpiredTimes = 600
)

var ctx = context.Background()

// GRedis struct
type GRedis struct {
	client     *redis.Client
	expiryTime int
}

// NewRedis new redis pointer
func NewRedis() *GRedis {
	redisConfig := config.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port),
		Password: redisConfig.Password,
		DB:       redisConfig.Database,
	})

	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		logger.Error(pong, err)
		return nil
	}
	expiryTime := config.Config.Cache.ExpiryTime
	if expiryTime <= 0 {
		expiryTime = RedisExpiredTimes
	}

	return &GRedis{client: rdb, expiryTime: expiryTime}
}

// IsConnected check redis is connected or not
func (g *GRedis) IsConnected() bool {
	if g.client == nil {
		return false
	}

	_, err := g.client.Ping(ctx).Result()
	if err != nil {
		return false
	}
	return true
}

// Get get key from redis
func (g *GRedis) Get(key string, data interface{}) error {
	val, err := g.client.Get(ctx, key).Bytes()
	if err == redis.Nil {
		return nil
	}
	if err != nil {
		logger.Info("Cache fail to get: ", err)
		return nil
	}
	logger.Debugf("Get from redis %s - %s", key, val)

	err = json.Unmarshal(val, &data)
	if err != nil {
		return err
	}

	return nil
}

// Set data to redis
func (g *GRedis) Set(key string, val []byte) error {
	err := g.client.Set(ctx, key, val, time.Duration(g.expiryTime)*time.Second).Err()
	if err != nil {
		logger.Error("Cache fail to set: ", err)
		return err
	}
	logger.Debugf("Set to redis %s - %s", key, val)

	return nil
}

// Remove return list keys from redis
func (g *GRedis) Remove(keys ...string) error {
	err := g.client.Del(ctx, keys...).Err()
	if err != nil {
		logger.Errorf("Cache fail to delete key %s: %s", keys, err)
		return err
	}
	logger.Debug("Cache deleted key", keys)

	return nil
}

// Keys get redis keys by pattern regex
func (g *GRedis) Keys(pattern string) ([]string, error) {
	keys, err := g.client.Keys(ctx, pattern).Result()
	if err != nil {
		return nil, err
	}

	return keys, nil
}
