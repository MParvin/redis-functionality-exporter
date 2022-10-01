/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	tools "github.com/mparvin/redis-functionality-exporter/tools"
	"github.com/spf13/cobra"
)

// checkRedisCmd represents the checkRedis command
var checkRedisCmd = &cobra.Command{
	Use:   "checkRedis",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		checkRedisFunctionality()
	},
}

func init() {
	rootCmd.AddCommand(checkRedisCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checkRedisCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checkRedisCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func checkRedisFunctionality() {
	config_data := tools.ReadConfig("")
	redis_server_ip := config_data["redis_server_ip"]
	redis_server_port := config_data["redis_server_port"]
	redis_server_password := config_data["redis_server_password"]
	redis_server_db := config_data["redis_server_db"]

	client := tools.ConnectRedis(redis_server_ip, redis_server_port, redis_server_password, redis_server_db)

	tools.WriteRedis("redis-functionality-exporter", "1", client)
	tools.ReadRedis("redis-functionality-exporter", "1", client)
	tools.DelRedis("redis-functionality-exporter", "1", client)
}
