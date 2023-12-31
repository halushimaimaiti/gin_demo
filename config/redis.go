package config

import "github.com/go-redis/redis/v8"

var Rdb *redis.Client // Redis 客户端全局变量

func InitRedis() {
	// 初始化 Redis 客户端
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis 地址
		Password: "",               // Redis 密码
		DB:       0,                // 默认数据库
	})
}
