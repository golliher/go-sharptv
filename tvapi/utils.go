package tvapi

import (
	"fmt"
	"net"

	"github.com/spf13/viper"
)

var ip, port string

func SendToTV(sharpCommand string, sharpParameter string) {
	cmdString := fmt.Sprintf("%4s%-4s\r", sharpCommand, sharpParameter)

	ip = viper.GetString("ip")
	port = viper.GetString("port")

	connect_string := fmt.Sprintf("%s:%s", ip, port)
	if viper.GetBool("debug") {
		fmt.Printf("Connecting to TV at %s\n", connect_string)
	}
	conn, err := net.Dial("tcp", connect_string)

	if err != nil {
		fmt.Println("Error connecting to TV.")
		return
	}

	if viper.GetBool("debug") {
		fmt.Printf("Sending command %v\n", cmdString)
	}

	fmt.Fprintf(conn, cmdString)
	if err != nil {
		fmt.Println("An error occured.")
		fmt.Println(err.Error())
	} else {
		if viper.GetBool("debug") {
			fmt.Printf(">>>> Sent %v\n", cmdString)
		}
	}

	tmp := make([]byte, 256)
	result, err := conn.Read(tmp)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		if viper.GetBool("debug") {
			fmt.Printf(">>>> Received: %s %s\n", tmp, string(result))
		}

	}
}
