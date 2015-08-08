package tvapi

import (
	"fmt"
	"testing"
)

// Posit: This is a fully qualified "unit" test

func Example_ParseResult() {
	result := []byte("asdfasdfsd\rasdfasdf")
	fmt.Println(parseResult(result))

	result = []byte("AeIoU\r00000000000000000000\r")
	fmt.Println(parseResult(result))

	// Output:
	// asdfasdfsd
	// AeIoU
}

// These are functional tests, right?
//  The succeed or fail depending on the status of the whole system.

func TestSendValidIP(t *testing.T) {
	result := Send("MUTE", "0001", "192.168.4.11", "10002")
	if result != "OK" {
		msg := fmt.Sprintf("Expected 'OK', got %s", result)
		t.Error(msg)
	}
}

func TestSendInvalidIP(t *testing.T) {
	result := Send("MUTE", "ON", "192.168.4.12", "10002")
	if result != "Error connecting to TV" {
		msg := fmt.Sprintf("Expected 'Error connecting to TV.', got '%s'", result)
		t.Error(msg)
	}
}

func TestSendInvalidCommand(t *testing.T) {
	result := Send("XXXX", "ON", "192.168.4.11", "10002")
	if result != "ERR" {
		msg := fmt.Sprintf("Expected ERR, got %s", result)
		t.Error(msg)
	}
}
