package main

import "github.com/spf13/cobra"
import "fmt"
import "os"
import "net"

// import "bufio"

func sendToTV(sharpCommand string, sharpParameter string) {

	debug := true
	cmdString := fmt.Sprintf("%4s%-4s\r", sharpCommand, sharpParameter)
	conn, err := net.Dial("tcp", "192.168.4.11:10002")

	if err != nil {
		fmt.Println("Error connecting to TV.")
		return
	}

	if debug {
		fmt.Printf("Sending command %v\n", cmdString)
	}

		fmt.Fprintf(conn, cmdString)
	if err != nil {
		fmt.Println("An error occured.")
		fmt.Println(err.Error())
	} else {
		if debug {
			fmt.Printf(">>>> Sent %v\n", cmdString)
		}
	}

	tmp := make([]byte, 256)
	result, err := conn.Read(tmp)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		if debug{
		fmt.Printf(">>>> Received: %s %s\n", tmp, string(result))
		}

	}

}

func main() {

	var sharptvCmd = &cobra.Command{
		Use:   "sharptv",
		Short: "sharptv is your command line interface to your television set",
		Long: `sharptv is the main command, used to control your TV

GoSharpTV is a hobbist project by an owner of a Sharp brand TV for other owners
of Sharp brand TVs.  It is implemented in the the Go programming lanugage.

`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("SharpTV command line remote control.")
		},
	}

	var cmdVolume = &cobra.Command{
		Use:   "volume [integer volume level to set]",
		Short: "Set the volume level of the TV.",
		Long: `Adjust the sound volume for the television.

You may find that a lower volume is more pleasant at night.
    `,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 1 {
				fmt.Printf("Setting volume to %v\n", args[0])
				sendToTV("VOLM", args[0])
			} else if len(args) == 0 {
				sendToTV("VOLM", "?")
			} else {
				cmd.Usage()
				os.Exit(1)
			}
		},
	}
	var cmdVolumeDown = &cobra.Command{
		Use:   "voldown",
		Short: "Sends a volume down button press",
		Long: `Adjust the sound volume for the television.

You may find that a lower volume is more pleasant at night.
    `,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("Reducing the volume")
				sendToTV("RCKY", "32")
			} else {
				cmd.Usage()
				os.Exit(1)
			}
		},
	}


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
				sendToTV("MUTE", "0000")

			case args[0] == "on":
				fmt.Println("Turning on mute.  This will silence the TV.")
				sendToTV("MUTE", "0001")

			case args[0] == "off":
				fmt.Println("Turning off mute.  This will return TV to the previous volume.")
				sendToTV("MUTE", "0002")

			case args[0] == "status":
				sendToTV("MUTE", "?")

			}
		},
	}

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
				sendToTV("POWR", "1")
			case args[0] == "off":
				fmt.Println("Turning off the TV.")
				sendToTV("POWR", "0")
			case args[0] == "status":
				sendToTV("POWR", "?")
			default:
				cmd.Usage()
			}
		},
	}

	var cmdInput = &cobra.Command{
		Use:   "input [TV source input number]",
		Short: "Set the input source",
		Long: `Adjust the input source to be displayed on the TV

    `,
		Run: func(cmd *cobra.Command, args []string) {
			switch {

			case len(args) == 1:
				fmt.Printf("Setting input source to %v\n", args[0])
				sendToTV("IAVD", args[0])

			case len(args) != 1:
				cmd.Usage()
			}
		},
	}
	sharptvCmd.AddCommand(cmdPower)
	sharptvCmd.AddCommand(cmdMute)
	sharptvCmd.AddCommand(cmdInput)
	sharptvCmd.AddCommand(cmdVolume)
	sharptvCmd.AddCommand(cmdVolumeDown)
	sharptvCmd.Execute()
}
