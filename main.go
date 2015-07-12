package main

import (
	"fmt"

	"github.com/golliher/go-sharptv/commands"
	"github.com/spf13/viper"
)

func main() {

	// Set config file
	viper.SetConfigName("config")

	// Add config path
	viper.AddConfigPath("$HOME/.sharptv")

	// Read in the config
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// Load default settings
	viper.SetDefault("debug", false)

	viper.SetEnvPrefix("gosharptv") // will be uppercased automatically
	viper.BindEnv("debug")
	viper.BindEnv("ip")
	viper.BindEnv("port")

	if viper.IsSet("ip") {
		ip := viper.GetString("ip")
		fmt.Printf("IP of TV to connect to: %s\n", ip)
	}

	// Do some flag handling and any complicated config logic

	// Start using configuration

	if viper.GetBool("debug") {
		fmt.Println("debug enabled")
	}

	commands.SharptvCmd.Execute()
}
