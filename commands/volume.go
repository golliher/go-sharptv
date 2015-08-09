package commands

import (
	"fmt"
	"os"
	"strconv"

	"github.com/golliher/go-sharptv/internal/github.com/spf13/cobra"
)

var cmdVolume = &cobra.Command{
	Use:   "volume {0..60|up|down|status}",
	Short: "Set the volume level of the TV.",
	Long: `Adjust the sound volume for the television.

You may find that a lower volume is more pleasant at night.

Examples:

	go-sharptv volume 0    # Effectively mutes without showing the mute icon.
	go-sharptv volume 25   # Set TV to a little less than half volume.
	go-sharptv volume 60   # Blast the volume as loud as it will go!
	go-sharptv volume down # Reduce the volume by one.
	go-sharptv status      # Returns information about the current volume level.


    `,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) != 1 {
			cmd.Usage()
			os.Exit(1)
		}

		InitializeConfig()

		numericalArgument, err := strconv.Atoi(args[0])
		if err == nil {
			if numericalArgument > -1 && numericalArgument < 61 {
				fmt.Printf("Setting volume to %v\n", args[0])
				sendToTV("VOLM", args[0])
			} else {
				fmt.Printf("Volume specificed is out of range 0 to 60")
				return
			}
			return
		}

		switch args[0] {

		case "down":
			fmt.Println("Reducing the volume")
			sendToTV("RCKY", "32")

		case "up":
			fmt.Println("Increasing the volume")
			sendToTV("RCKY", "33")

		case "status":
			result := sendToTV("VOLM", "?")
			if result != "ERR" {
				fmt.Printf("Volume is: %v\n", result)
			} else {
				fmt.Println("Unable to determine current volume.")
			}

		default:
			cmd.Usage()
		}

	},
}

func init() {
	SharptvCmd.AddCommand(cmdVolume)
}
