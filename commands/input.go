package commands

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golliher/go-sharptv/internal/github.com/spf13/cobra"
)

var cmdInput = &cobra.Command{
	Use:   "input {1-8|tv|status}",
	Short: "Set the input source",
	Long: `Adjust the input source to be displayed on the TV

	examples:
		go-sharptv input 4    # switches to the 4th input source
		go-sharptv input tv   # switches to the TV tuner
		go-sharptv input status # Returns information on the current input source

  `,
	Run: func(cmd *cobra.Command, args []string) {

		switch {

		case len(args) == 1:
			switch args[0] {

			case "status":
				result := sendToTV("IAVD", "?")
				if result != "ERR" {

					// Find and print the friendly name
					for key := range inputLabelMap {
						s, _ := strconv.Atoi(result)
						if inputLabelMap[key] == s {
							fmt.Printf("Input is: %s (%s)\n", result, key)
							return
						}
					}
					fmt.Printf("Input is: %s\n", result)

				}
			case "tv":
				result := sendToTV("ITVD", "0")
				if result == "OK" {
					fmt.Println("Switched input TV")
				} else {
					fmt.Println("Unable to switch to TV")
					os.Exit(1)
				}
			case "1", "2", "3", "4", "5", "6", "7", "8", "9":
				result := sendToTV("IAVD", args[0])
				if result == "OK" {
					fmt.Printf("Input source changed to %v\n", args[0])
				} else {
					fmt.Printf("Unable to set source to %s.  '%s' is not valid.\n", args[0], args[0])
					os.Exit(1)
				}
			default:

				// Finally we check to see if we have a match in the
				// configured inputLabelMap of friendly names to input numbers
				if inputLabelMap[args[0]] != 0 {

					s := strconv.Itoa(inputLabelMap[args[0]])
					result := sendToTV("IAVD", s)
					if result == "OK" {
						fmt.Printf("Input source changed to %v\n", args[0])
					} else {
						// A possible reason is that is already the source
						// To be extra helpful we check to see if that is the
						// case before decided to exit with success for fail.

						// BUG(golliher):  This fails likley a timing issue with
						// trying to reconnect so quickly, Hack is to sleep
						// 10 Milliseconds seems to be enough on my system.
						time.Sleep(10 * time.Millisecond)
						result = sendToTV("IAVD", "?")
						if result == s {
							fmt.Println("That is already the active source.")
							os.Exit(0)
						}
						fmt.Printf("Unable to set source to %s.\n", args[0])
						os.Exit(1)
					}

				} else {
					// We've tried everything to match the users source request
					// It's time to give up.
					fmt.Printf("Unable to set source to %s.  '%s' is not recognized.\n", args[0], args[0])
					os.Exit(1)
				}

			}

		case len(args) != 1:
			cmd.Usage()
		}
	},
}

func init() {
	SharptvCmd.AddCommand(cmdInput)
}
