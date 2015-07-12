package main

import (
	"fmt"
	"os"

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

	// Do some flag handling and any complicated config logic
	if !viper.IsSet("ip") || !viper.IsSet("port") {
		fmt.Println("Configuration error.  Both IP and PORT must be set via either config, environment, or flags.")
		os.Exit(1)
	}

	// Setup global flags
	commands.SharptvCmd.PersistentFlags().BoolP("debug", "d", false, "Print debug messages")
	viper.BindPFlag("debug", commands.SharptvCmd.PersistentFlags().Lookup("debug"))

	// Start using configuration

	if viper.GetBool("debug") {
		fmt.Println("debug enabled")
	}

	commands.SharptvCmd.Execute()
}
