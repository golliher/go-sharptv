package commands

import (
	"fmt"

	"github.com/golliher/go-sharptv/tvapi"
	"github.com/spf13/cobra"
)

var cmdInput = &cobra.Command{
	Use:   "input [TV source input number]",
	Short: "Set the input source",
	Long: `Adjust the input source to be displayed on the TV

  `,
	Run: func(cmd *cobra.Command, args []string) {
		switch {

		case len(args) == 1:
			fmt.Printf("Setting input source to %v\n", args[0])
			tvapi.SendToTV("IAVD", args[0])

		case len(args) != 1:
			cmd.Usage()
		}
	},
}

func init() {
	SharptvCmd.AddCommand(cmdInput)
}
