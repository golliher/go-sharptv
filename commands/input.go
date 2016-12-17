package commands

import (
	"fmt"
	"os"
	"strconv"

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
				input, err := tv.GetInput()
				checkErr(err)

				// Find and print the friendly name
				for friendly_name := range inputLabelMap {
					if strconv.Itoa(inputLabelMap[friendly_name]) == input {
						fmt.Printf("Input is: %d (%s)\n",
							input, friendly_name)
						return
					}
				}
				fmt.Printf("Input is: %d\n", input)

			case "tv":
				err := tv.SetInput("0")
				if err != nil {
					fmt.Println("Switched input TV")
				} else {
					fmt.Println("Unable to switch to TV")
					os.Exit(1)
				}
			case "1", "2", "3", "4", "5", "6", "7", "8", "9":

				input, err := tv.GetInput()
				checkErr(err)

				if input == args[0] {
					fmt.Println("That is already the active source.")
					os.Exit(0)
				}

				err = tv.SetInput(args[0])
				checkErr(err)
				fmt.Println("Input source changed to %v\n", args[0])
			default:

				// Finally we check to see if we have a match in the
				// configured inputLabelMap of friendly names to input numbers
				if inputLabelMap[args[0]] != 0 {

					s := strconv.Itoa(inputLabelMap[args[0]])

					input, err := tv.GetInput()
					checkErr(err)

					if input == s {
						fmt.Println("That is already the active source.")
						os.Exit(0)
					}

					err = tv.SetInput(s)
					checkErr(err)
					fmt.Println("Input source changed to %v\n", args[0])

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
