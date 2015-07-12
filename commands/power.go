package commands

import (
	"fmt"

	"github.com/golliher/go-sharptv/tvapi"
	"github.com/spf13/cobra"
)

var cmdPower = &cobra.Command{
	Use:   "power {on|off}",
	Short: "Turn the TV off or on",
	Long: `Powers the TV off or on.  If neither subcommand of either "off" nor "on" are
  specfified, then the power will be toggled from it's current state.`,
	Run: func(cmd *cobra.Command, args []string) {

		switch {
		case len(args) == 0:
			fmt.Println("Toggling Power is not yet implemented.")
		case len(args) > 1:
			cmd.Usage()
		case args[0] == "on":
			fmt.Println("Turning on the TV.")
			tvapi.SendToTV("POWR", "1")
		case args[0] == "off":
			fmt.Println("Turning off the TV.")
			tvapi.SendToTV("POWR", "0")

		case args[0] == "status":

			result := tvapi.SendToTV("POWR", "?")

			switch result {
			case "1":
				fmt.Println("TV is ON")
			case "0":
				fmt.Println("TV is OFF")
			default:
				fmt.Printf("Warning: unexpected result >%v<\n\n", result)
			}

		default:
			cmd.Usage()
		}
	},
}

func init() {
	SharptvCmd.AddCommand(cmdPower)
}
