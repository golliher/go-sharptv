package commands

import (
	"fmt"

	"github.com/golliher/go-sharptv/tvapi"

	"github.com/golliher/go-sharptv/internal/github.com/spf13/cobra"
)

var cmdMute = &cobra.Command{
	Use:   "mute {on|off}",
	Short: "Turn the volume of the TV off or on",
	Long: `Mutes or unmutes the television.  If not subcommand of either "off" or "on" are
  specfified, then the mute will be toggled from it's current state.`,
	Run: func(cmd *cobra.Command, args []string) {
		switch {
		case len(args) > 1:
			cmd.Usage()
		case len(args) == 0:
			fmt.Println("Toggling mute.")
			tvapi.SendToTV("MUTE", "0000")

		case args[0] == "on":
			fmt.Println("Turning on mute.  This will silence the TV.")
			tvapi.SendToTV("MUTE", "0001")

		case args[0] == "off":
			fmt.Println("Turning off mute.  This will return TV to the previous volume.")
			tvapi.SendToTV("MUTE", "0002")

		case args[0] == "status":
			result := tvapi.SendToTV("MUTE", "?")
			switch result {
			case "1":
				fmt.Println("TV is muted")
			case "2":
				fmt.Println("TV is not muted")
			default:
				fmt.Printf("Warning: unexpected result >%v<\n\n", result)
			}

		}
	},
}

func init() {
	SharptvCmd.AddCommand(cmdMute)
}
