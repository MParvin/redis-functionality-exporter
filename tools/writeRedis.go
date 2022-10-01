package tools

import (
	"fmt"

	"github.com/go-redis/redis"
)

func WriteRedis(key string, value string, client *redis.Client) bool {
	config_data := ReadConfig("")
	redis_server_ip := config_data["redis_server_ip"]
	redis_server_port := config_data["redis_server_port"]
	redis_server_password := config_data["redis_server_password"]
	redis_server_db := config_data["redis_server_db"]

	// Create a redis client
	ConnectRedis(redis_server_ip, redis_server_port, redis_server_password, redis_server_db)

	err := client.Set(key, value, 0).Err()
	if err != nil {
		fmt.Println("Could not write to redis-server, ", err)
		return false
	}
	fmt.Println("Successfully wrote to redis-server")
	return true
}
