package tools

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
)

// ConnectRedis will connect to redis-server and return a redis client
func ConnectRedis(redis_server_ip string, redis_server_port string, redis_server_password string, redis_server_db string) *redis.Client {
	// Create a redis client
	client := redis.NewClient(&redis.Options{
		Addr:        redis_server_ip + ":" + redis_server_port,
		Password:    redis_server_password,
		DB:          0,
		DialTimeout: 6,
	})
	// Check if redis-server is running
	_, err := client.Ping().Result()
	if err != nil {
		fmt.Println("Error connecting to redis-server, ", err)
		os.Exit(1)
	}
	return client
}

func CheckRedisConnection(client *redis.Client) bool {
	_, err := client.Ping().Result()
	if err != nil {
		return false
	}
	return true
}
