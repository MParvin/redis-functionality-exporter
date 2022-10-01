/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	tools "github.com/mparvin/redis-functionality-exporter/tools"
	"github.com/spf13/cobra"
)

func showHelp() {
	fmt.Println("redis-fe config generate")
}

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Generate a config file for redis-functionality-exporter",
	Long: `Generate a config file for redis-functionality-exporter
	You can use the following command to generate a config file:
	redis-fe config generate`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("You need to specify a subcommand")
			showHelp()
			os.Exit(1)
		}
		if args[0] == "generate" {
			config_file := "config.yml"
			// if -f flag is used, then use the value of -f flag as config file name
			if cmd.Flags().Changed("file") {
				config_file, _ = cmd.Flags().GetString("file")
			}
			generate_config(config_file)
		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	configCmd.PersistentFlags().String("generate", "", "Generate a config file for redis-functionality-exporter")

	// -f flag is used to specify a config file name
	configCmd.Flags().StringP("file", "f", "config.yml", "use -f to specify a config file name")
}

// func generate_config will generate a simple config file that includes
// redis-server ip address
// redis-server password
// redis-server port
// redis-server database is optional
func generate_config(config_file string) {
	// Ask user to enter redis-server ip address
	var redis_server_ip string
	fmt.Println("Enter redis-server ip address: ")
	fmt.Scanln(&redis_server_ip)
	// Validate redis-server ip address
	if redis_server_ip == "" {
		fmt.Println("redis-server ip address is required")
		os.Exit(1)
	}

	if !tools.ValidateIP(redis_server_ip) {
		fmt.Println("Invalid redis-server ip address")
		os.Exit(1)
	}

	var redis_server_password string
	fmt.Println("Enter redis-server password: ")
	fmt.Scanln(&redis_server_password)

	var redis_server_port string
	fmt.Println("Enter redis-server port: [default: 6379]")
	fmt.Scanln(&redis_server_port)
	// If user enter empty string for redis-server port, then use default port 6379
	if redis_server_port == "" {
		redis_server_port = "6379"
	}

	var redis_server_database string
	fmt.Println("Enter redis-server database: [optional]")
	fmt.Scanln(&redis_server_database)
	// If user enter empty string for redis-server database, then use default database 0
	if redis_server_database == "" {
		redis_server_database = "0"
	}
	// Write data to config.yml file
	writeConfigToFile(config_file, redis_server_ip, redis_server_password, redis_server_port, redis_server_database)

}

func writeConfigToFile(config_file string, redis_server_ip string, redis_server_password string, redis_server_port string, redis_server_database string) {
	var err error
	var file *os.File
	// If config file exists, then delete it
	if _, err = os.Stat(config_file); err == nil {
		err = os.Remove(config_file)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
	}
	// Create config file
	file, err = os.Create(config_file)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer file.Close()
	file.WriteString("redis_server_ip: " + redis_server_ip + "\n")
	file.WriteString("redis_server_password: " + redis_server_password + "\n")
	file.WriteString("redis_server_port: " + redis_server_port + "\n")
	file.WriteString("redis_server_database: " + redis_server_database + "\n")
}
