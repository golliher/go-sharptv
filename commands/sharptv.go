package commands

import (
	"fmt"

	"github.com/golliher/go-sharptv/internal/github.com/spf13/cobra"
)

var SharptvCmd = &cobra.Command{
	Use:   "sharptv",
	Short: "sharptv is your command line interface to your television set",
	Long: `sharptv is the main command, used to control your TV

GoSharpTV is a hobbist project by an owner of a Sharp brand TV for other owners
of Sharp brand TVs.  It is implemented in the the Go programming lanugage.

`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("SharpTV command line remote control.")
	},
}
