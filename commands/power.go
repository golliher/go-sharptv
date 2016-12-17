package commands

import (
	"fmt"

	"github.com/golliher/go-sharptv/internal/github.com/spf13/cobra"
)

var cmdOff = &cobra.Command{
	Use:   "off",
	Short: "Turn the TV off",
	Long:  `Powers the TV off.  It is equivalent to the command "power off".`,
	Run: func(cmd *cobra.Command, args []string) {

		switch {
		case len(args) == 0:
			{
				err := tv.PowerOff()
				checkErr(err)
			}
		default:
			cmd.Usage()
		}
	},
}

var cmdOn = &cobra.Command{
	Use:   "on",
	Short: "Turn the TV on",
	Long:  `Powers the TV on.  It is equivalent to the command "power on".`,
	Run: func(cmd *cobra.Command, args []string) {
		switch {
		case len(args) == 0:
			{
				err := tv.PowerOn()
				checkErr(err)
			}
		default:
			cmd.Usage()
		}
	},
}

var cmdPower = &cobra.Command{
	Use:   "power {on|off|status}",
	Short: "Turn the TV off or on",
	Long: `Powers the TV off or on.  If neither subcommand of either "off" nor "on" are
  specfified, then the power will be toggled from it's current state.   If
	"status" is specified, returns information for the current power state.`,
	Run: func(cmd *cobra.Command, args []string) {

		switch {
		case len(args) == 0:
			{
				// Toggle power if no argument given to power command

				if tv.IsOn() {
					err := tv.PowerOff()
					checkErr(err)
				} else {
					err := tv.PowerOn()
					checkErr(err)
				}
			}
		case len(args) > 1:
			cmd.Usage()
		case args[0] == "on":
			if tv.IsOn() {
				fmt.Println("TV is already on")
				return
			}
			fmt.Println("Turning on the TV.")
			err := tv.PowerOn()
			checkErr(err)
		case args[0] == "off":
			if tv.IsOff() {
				fmt.Println("TV is already off")
				return
			}
			fmt.Println("Turning off the TV.")
			err := tv.PowerOff()
			checkErr(err)
		case args[0] == "status":
			switch tv.IsOn() {
			case true:
				fmt.Println("TV is ON")
			case false:
				fmt.Println("TV is OFF")
			}
		default:
			cmd.Usage()
		}
	},
}

func init() {
	SharptvCmd.AddCommand(cmdPower)
	SharptvCmd.AddCommand(cmdOn)
	SharptvCmd.AddCommand(cmdOff)
}
