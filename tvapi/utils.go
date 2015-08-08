package tvapi

import (
	"fmt"
	"net"
	"strings"

	"github.com/golliher/go-sharptv/internal/github.com/spf13/viper"
)

var ip, port string

// Pull out the characters up to the first \r
func parseResult(resultstring []byte) string {
	parsed := strings.Split(string(resultstring), "\r")
	return parsed[0]
}

// SendToTV transmits Sharp Aquos API commands to the Television over the network
func SendToTV(sharpCommand string, sharpParameter string) string {
	cmdString := fmt.Sprintf("%4s%-4s\r", sharpCommand, sharpParameter)

	ip = viper.GetString("ip")
	port = viper.GetString("port")

	connectString := fmt.Sprintf("%s:%s", ip, port)
	if viper.GetBool("debug") {
		fmt.Printf("Connecting to TV at %s\n", connectString)
	}
	conn, err := net.Dial("tcp", connectString)

	if err != nil {
		fmt.Println("Error connecting to TV.")
		return ("Error connecting to TV")
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

	apiResult := make([]byte, 32)
	bytesRead, err := conn.Read(apiResult)
	if err != nil {
		fmt.Println(err.Error())
	} else {

		resultString := parseResult(apiResult)

		if viper.GetBool("debug") {
			fmt.Printf(">>>> Received: %s Bytes: %v\n", resultString, bytesRead)
		}
		conn.Close()
		return resultString
	}
	return "no result"
}
