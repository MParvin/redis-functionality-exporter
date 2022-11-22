/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	tools "github.com/mparvin/redis-functionality-exporter/tools"
	"github.com/spf13/cobra"
)

// checkRedisCmd represents the checkRedis command
var checkRedisCmd = &cobra.Command{
	Use:   "checkRedis",
	Short: "Check redis-server functionality",
	Long: `This command will check redis-server functionality and print the output on the screen.
	You can use the following command to check redis-server functionality:
	redis-fe checkRedis

	To use it as a prometheus exporter, you can use the following command:
	redis-fe checkRedis --prometheus`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(tools.CheckRedisFunctionality())
	},
}

func init() {
	rootCmd.AddCommand(checkRedisCmd)
	// If -p flag is not provided, it will use the default value of 8080
	checkRedisCmd.Flags().IntP("port", "p", 8080, "Port to run the exporter on")

	// handle --prometheus flag
	checkRedisCmd.Flags().Bool("prometheus", false, "Run the exporter as a prometheus exporter")

	// Call RunWebServer with the port value
	checkRedisCmd.Run = func(cmd *cobra.Command, args []string) {
		port, _ := cmd.Flags().GetInt("port")
		prometheus, _ := cmd.Flags().GetBool("prometheus")
		if prometheus {
			tools.RunWebServer(port)
		} else {
			fmt.Println(tools.CheckRedisFunctionality())
		}
	}
}
