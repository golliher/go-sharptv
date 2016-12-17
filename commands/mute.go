package commands

import (
	"fmt"
	"time"

	"github.com/golliher/go-sharptv/internal/github.com/spf13/cobra"
)

var cmdMute = &cobra.Command{
	Use:   "mute {on|off|status}",
	Short: "Turn the volume of the TV off or on",
	Long: `Mutes or unmutes the television.  If not subcommand of either "off" or "on" are
  specfified, then the mute will be toggled from it's current state.  If "status" is
	specified, then information will be returned about the curret state of the mute function.`,
	Run: func(cmd *cobra.Command, args []string) {

		switch {
		case len(args) > 1:
			cmd.Usage()
		case len(args) == 0:
			err := tv.ToggleMute()
			checkErr(err)
			time.Sleep(100 * time.Millisecond)
			if tv.IsMuted() {
				fmt.Println("Mute toggled to MUTED")
			} else {
				fmt.Println("Mute toggled to UNMUTED")
			}

		case args[0] == "on":
			err := tv.Mute()
			checkErr(err)
			fmt.Println("Turning on mute. This will silence the TV and remeber the volume.")

		case args[0] == "off":
			err := tv.Unmute()
			checkErr(err)
			fmt.Println("Unmuted. This returns TV to the previous volume.")

		case args[0] == "status":
			if !tv.IsOn() {
				fmt.Println("TV is OFF. Mute has no effect.")
				return
			}
			switch tv.IsMuted() {
			case true:
				fmt.Println("TV is muted")
			case false:
				fmt.Println("TV is not muted")
			}
		}
	},
}

func init() {
	SharptvCmd.AddCommand(cmdMute)
}
