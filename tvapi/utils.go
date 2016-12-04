package tvapi

import (
	"fmt"
	"net"
	"strings"
	"time"
)

type TV struct {
	IP   string
	Port string
}

func (tv TV) PowerOff() error {
	result := tv.Send("POWR", "0")
	switch result {
	case "ERR":
		fmt.Println("Something went wrong.  Attempted to turn TV on and failed.")
	case "OK":
		return nil
	default:
		fmt.Printf("Warning: unexpected result >%v<\n\n", result)
	}

	return nil
}

func (tv TV) PowerOn() error {
	result := tv.Send("POWR", "1")
	switch result {
	case "ERR":
		fmt.Println("Something went wrong.  Attempted to turn TV on and failed.")
	case "OK":
		return nil
	default:
		fmt.Printf("Warning: unexpected result >%v<\n\n", result)
	}

	return nil
}

// Pull out the characters up to the first \r
func parseResult(resultstring []byte) string {
	parsed := strings.Split(string(resultstring), "\r")
	return parsed[0]
}

// Send transmits Sharp Aquos API commands to the Television over the network
func (tv TV) Send(sharpCommand string, sharpParameter string) string {
	cmdString := fmt.Sprintf("%4s%-4s\r", sharpCommand, sharpParameter)

	connectString := fmt.Sprintf("%s:%s", tv.IP, tv.Port)
	conn, err := net.DialTimeout("tcp", connectString, time.Duration(100*time.Millisecond))

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
		fmt.Printf("Only read in %d bytes:", bytesRead)

	} else {
		resultString := parseResult(apiResult)
		conn.Close()
		return resultString
	}

	return "no result"
}
