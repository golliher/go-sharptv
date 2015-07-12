package main

import (
	"fmt"

	"github.com/golliher/go-sharptv/commands"
	"github.com/spf13/viper"
)

// Flags that are to be added to commands
// var ip, port string

func main() {

	viper.SetConfigName("config")
	viper.AddConfigPath("$HOME/.sharptv")
	viper.SetDefault("debug", false)

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	if viper.GetBool("debug") {
		fmt.Println("debug enabled")
	}

	commands.SharptvCmd.Execute()
}
