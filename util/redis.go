package util

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

var RedisClient *redis.Pool
var RedisUtil *redisUtil

func init() {
	RedisClient = &redis.Pool{
		// 从配置文件获取maxidle以及maxactive，取不到则用后面的默认值
		MaxIdle:     2,
		MaxActive:   10,
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "qq10086.zhidaohu.com:6389")
			if err != nil {
				return nil, err
			}

			// 选择db
			c.Do("SELECT", 0)
			return c, nil
		},
	}

	RedisUtil = &redisUtil{}
}

type redisUtil struct {
}

func (*redisUtil) Get(key string) (string, error) {
	client := RedisClient.Get()
	defer client.Close()
	return redis.String(client.Do("GET", key))
}

func (*redisUtil) Set(key string, value string, timeout int) (string, error) {
	client := RedisClient.Get()
	defer client.Close()
	return redis.String(client.Do("SET", key, value, "EX", timeout))
}
