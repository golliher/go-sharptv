// A command line Aquos Remote Control for Sharp brand televisions.
// Certain television sets made by Sharp Inc. respond to commands over
// a simple network protocol.   This enables applications to control the
// TV in the same manner than an IR remote control can.
//
// This package implements a CLI interface to those televisions.
package main

import "github.com/golliher/go-sharptv/commands"

func main() {
	commands.SharptvCmd.Execute()
}
