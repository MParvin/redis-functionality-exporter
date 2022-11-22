package tools

func CheckRedisFunctionality() string {
	config_data := ReadConfig("")
	redis_server_ip := config_data["redis_server_ip"]
	redis_server_port := config_data["redis_server_port"]
	redis_server_password := config_data["redis_server_password"]
	redis_server_db := config_data["redis_server_db"]

	client, err := ConnectRedis(redis_server_ip, redis_server_port, redis_server_password, redis_server_db)
	if err != nil {
		return "redis_connection 0"
	}

	write_status := WriteRedis("redis-functionality-exporter", "1", client)
	read_status := ReadRedis("redis-functionality-exporter", "1", client)
	delete_status := DelRedis("redis-functionality-exporter", "1", client)

	connection_status := "redis_connection 1"
	write_status_string := "redis_write 0"
	read_status_string := "redis_read 0"
	delete_status_string := "redis_delete 0"

	return_value := "redis_connection 1"
	if write_status {
		write_status_string = "redis_write 1"
	}

	if read_status {
		read_status_string = "redis_read 1"
	}

	if delete_status {
		delete_status_string = "redis_delete 1"
	}

	return_value = connection_status + "\n" + write_status_string + "\n" + read_status_string + "\n" + delete_status_string
	return return_value
}
