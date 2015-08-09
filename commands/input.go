package commands

import (
	"fmt"

	"github.com/golliher/go-sharptv/internal/github.com/spf13/cobra"
)

var cmdInput = &cobra.Command{
	Use:   "input [TV source input number]",
	Short: "Set the input source",
	Long: `Adjust the input source to be displayed on the TV

  `,
	Run: func(cmd *cobra.Command, args []string) {

		InitializeConfig()

		switch {

		case len(args) == 1:
			switch args[0] {

			case "status":
				result := sendToTV("IAVD", "?")
				if result != "ERR" {
					fmt.Printf("Input is: %s\n", result)
				}
			case "tv":
				result := sendToTV("ITVD", "0")
				if result == "OK" {
					fmt.Println("Switched input TV")
				} else {
					fmt.Println("Unable to switch to TV")
				}
			default:
				fmt.Printf("Setting input source to %v\n", args[0])
				sendToTV("IAVD", args[0])

			}

		case len(args) != 1:
			cmd.Usage()
		}
	},
}

func init() {
	SharptvCmd.AddCommand(cmdInput)
}
