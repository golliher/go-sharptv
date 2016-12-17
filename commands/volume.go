package commands

import (
	"fmt"
	"os"
	"strconv"
	"time"

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

		// If we can convert to integer, assume user is setting volume to value
		_, err := strconv.Atoi(args[0])
		if err == nil {
			err := tv.SetVolume(args[0])
			checkErr(err)
			os.Exit(0)
		}

		// Otherwise user has given a sub-command string
		switch args[0] {

		case "down":
			err := tv.DecreaseVolume()
			checkErr(err)
			time.Sleep(10 * time.Millisecond)
			volume, _ := tv.GetVolume()
			fmt.Printf("Volume decreased to %s\n", volume)

		case "up":
			err := tv.IncreaseVolume()
			checkErr(err)
			time.Sleep(10 * time.Millisecond)
			volume, _ := tv.GetVolume()
			fmt.Printf("Volume increased to %s\n", volume)

		case "status":
			result, _ := tv.GetVolume()
			if result != "ERR" {
				fmt.Printf("Volume is: %v\n", result)
			} else {
				fmt.Println("Unable to determine current volume. Is the TV on?")
				os.Exit(1)
			}

		default:
			cmd.Usage()
		}
	},
}

func init() {
	SharptvCmd.AddCommand(cmdVolume)
}
