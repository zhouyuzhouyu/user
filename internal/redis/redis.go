package redis

import (
    "github.com/go-redis/redis/v8"
)

var RedisDB *redis.Client

// func init() {
// 	redisEntry := rkredis.GetRedisEntry("redis")
// 	RedisClient, _ = redisEntry.GetClient()
// }

func SetupRedis(redisClien *redis.Client) {
    RedisDB = redisClien
}
