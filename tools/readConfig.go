package tools

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// ReadConfig will read config file and return a map of config values
func ReadConfig(config_file string) map[string]string {
	if config_file == "" {
		config_file = "config.yml"
	}
	// Read config file
	viper.SetConfigFile(config_file)
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading config file, %s", err)
		os.Exit(1)
	}
	// Get all keys from config file
	keys := viper.AllKeys()
	// Create a map to store config values
	config := make(map[string]string)
	// Iterate over all keys and get value for each key
	for _, key := range keys {
		value := viper.Get(key)
		// Convert value to string
		config[key] = fmt.Sprintf("%v", value)
	}
	return config
}
