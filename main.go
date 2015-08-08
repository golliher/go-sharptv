// A command line Aquos Remote Control for Sharp brand televisions.
// Certain television sets made by Sharp Inc. respond to commands over
// a simple network protocol.   This enables applications to control the
// TV in the same manner than an IR remote control can.
//
// This package implements a CLI interface to those televisions.
package main

import (
	"fmt"
	"os"

	"github.com/golliher/go-sharptv/commands"
	"github.com/golliher/go-sharptv/internal/github.com/spf13/viper"
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
