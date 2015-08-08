package tvapi

import (
	"fmt"
	"net"
	"strings"
	"time"
)

// Pull out the characters up to the first \r
func parseResult(resultstring []byte) string {
	parsed := strings.Split(string(resultstring), "\r")
	return parsed[0]
}

// Send transmits Sharp Aquos API commands to the Television over the network
func Send(sharpCommand string, sharpParameter string, ip string, port string) string {
	cmdString := fmt.Sprintf("%4s%-4s\r", sharpCommand, sharpParameter)

	connectString := fmt.Sprintf("%s:%s", ip, port)
	conn, err := net.DialTimeout("tcp", connectString, time.Duration(10*time.Millisecond))

	if err != nil {
		fmt.Println("Error connecting to TV.")
		return ("Error connecting to TV")
	}

	fmt.Fprintf(conn, cmdString)
	if err != nil {
		fmt.Println("An error occured.")
		fmt.Println(err.Error())
	}

	apiResult := make([]byte, 32)
	bytesRead, err := conn.Read(apiResult)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Printf("Only red in %d bytes:", bytesRead)

	} else {
		resultString := parseResult(apiResult)
		conn.Close()
		return resultString
	}

	return "no result"
}
