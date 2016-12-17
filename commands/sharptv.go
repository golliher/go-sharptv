package commands

import (
	"fmt"
	"os"

	"github.com/golliher/go-sharptv/internal/github.com/spf13/cobra"
	"github.com/golliher/go-sharptv/tvapi"
	"github.com/spf13/viper"
)

var tv tvapi.TV

var inputLabelMap map[string]int // map of input labels e.g. inputLabels[1] == ""

// SharptvCmd is the root command
var SharptvCmd = &cobra.Command{
	Use:   "sharptv",
	Short: "sharptv is your command line interface to your television set",
	Long: `sharptv is the main command, used to control your TV

Go-SharpTV is a hobbist project by an owner of a Sharp brand TV for other owners
of Sharp brand TVs.  It is implemented in the the Go programming lanugage.

`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("SharpTV command line remote control.")
	},
}

// InitializeConfig loads our configuration using Viper package.
func InitializeConfig() {
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

	// TODO --implement the use of this data in the input command

	inputLabelMap = make(map[string]int)

	inputNames := []string{"input1", "input2", "input3", "input4", "input5", "input6", "input7", "input8"}
	for i, v := range inputNames {
		if viper.IsSet(v) {
			inputname := viper.GetString(v)
			inputLabelMap[inputname] = i + 1
		}
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}
}

func init() {
	InitializeConfig()
	tv.IP = viper.GetString("ip")
	tv.Port = viper.GetString("port")
}
