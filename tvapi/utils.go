package tvapi

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

type TV struct {
	IP   string
	Port string
}

// Functions that make changes
func (tv TV) Mute() error {

	if tv.IsOff() {
		return errors.New("Mute command is ineffecitve while TV is off")
	}
	_, err := tv.send("MUTE", "0001")
	if err != nil {
		return fmt.Errorf("Error attepting to mute: %s", err)
	}
	return nil
}
func (tv TV) Unmute() error {
	if tv.IsOff() {
		return errors.New("Unmute command is ineffecitve while TV is off")
	}
	_, err := tv.send("MUTE", "0002")
	return err
}
func (tv TV) ToggleMute() error {
	if tv.IsOff() {
		return errors.New("Mute toggle is ineffecitve while TV is off")
	}
	_, err := tv.send("MUTE", "0000")
	return err
}
func (tv TV) PowerOff() error {
	_, err := tv.send("POWR", "0")
	return err
}
func (tv TV) PowerOn() error {
	_, err := tv.send("POWR", "1")
	return err
}
func (tv TV) SetInput(newinput string) error {
	_, err := tv.send("IAVD", newinput)
	if err != nil {
		return fmt.Errorf("ERROR setting input (%d): %s", newinput, err)
	}
	return err
}

func (tv TV) IncreaseVolume() error {
	_, err := tv.send("RCKY", "33")
	return err
}

func (tv TV) DecreaseVolume() error {
	_, err := tv.send("RCKY", "32")
	return err
}

func (tv TV) SetVolume(newvolume string) error {

	numericalArgument, err := strconv.Atoi(newvolume)
	if err != nil {
		return fmt.Errorf("Invalid volume %s", newvolume)
	}

	if numericalArgument < 1 || numericalArgument > 60 {
		return fmt.Errorf("Volume specified (%s )is out of range 0 to 60", newvolume)
	}

	_, err = tv.send("VOLM", newvolume)
	if err != nil {
		return fmt.Errorf("Unable to set volume to %s. Is the TV on?", newvolume)
	}
	return err
}

// Functions that give status
func (tv TV) IsMuted() bool {

	result, err := tv.send("MUTE", "?")
	if err != nil {
		return false
	}

	switch result {
	case "1":
		return true
	case "2":
		return false
	default:
		return false
	}
}

func (tv TV) IsOff() bool {
	return !tv.IsOn()
}

func (tv TV) IsOn() bool {
	result, err := tv.send("POWR", "?")
	if err != nil {
		return false
	}
	if result == "1" {
		return true
	}
	return false
}

// Consider API change to just return string and not err
// potentially retry 3 times then just return UNKNOWN
func (tv TV) GetInput() (string, error) {
	return tv.send("IAVD", "?")
}

func (tv TV) GetVolume() (string, error) {
	return tv.send("VOLM", "?")
}

// Pull out the characters up to the first \r
func parseResult(resultstring []byte) string {
	parsed := strings.Split(string(resultstring), "\r")
	return parsed[0]
}

// Send transmits Sharp Aquos API commands to the Television over the network
func (tv TV) send(sharpCommand string, sharpParameter string) (string, error) {
	cmdString := fmt.Sprintf("%4s%-4s\r", sharpCommand, sharpParameter)

	connectString := fmt.Sprintf("%s:%s", tv.IP, tv.Port)
	conn, err := net.DialTimeout("tcp", connectString, time.Duration(100*time.Millisecond))
	if err != nil {
		return "", fmt.Errorf("Error connecting to TV: %s", err)
	}
	defer conn.Close()

	fmt.Fprintf(conn, cmdString)
	if err != nil {
		return "", fmt.Errorf("Error sending command to TV: %s, err")
	}

	apiResult := make([]byte, 32)
	bytesRead, err := conn.Read(apiResult)
	if err != nil {
		return "", fmt.Errorf("Error reading response from TV: Only read in %d bytes:",
			bytesRead)
	} else {
		resultString := parseResult(apiResult)
		if resultString == "ERR" {
			return resultString, errors.New("Error(ERR) returned by TV in response to command.")
		}
		return resultString, nil
	}

	// Can we even get here?  Hmm...
	return "", errors.New("BUG: Send() in utils.go fell through to the end.  That's not supposed to happen.")
}
