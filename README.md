# Redis functionality exporter

This is a simple exporter for Redis that exports the following metrics:

* redis_connection_time (in seconds)
* is_working (1 if the exporter is working, 0 if not) - it uses checkRedisFunctionality() to check is the redis server is working
* redis_write_time (in seconds) - it uses WriteRedis() to check the write time
* redis_read_time (in seconds) - it uses ReadRedis() to check the read time
* redis_delete_time  (in seconds) - it uses DeleteRedis() to check the delete time

TODO:
- [] replace `fmt.Println` and `os.Exit` with `log.Fatal`
- [] read configuration from environment variables and if it's not set read it from `config.yml` file
- [] write deployment
- [] create configmap example
- [] write test
- [] update readme