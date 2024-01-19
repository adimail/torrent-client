// main_test.go
package main

import (
	"bytes"
	"testing"
)

func TestPrintToConsole(t *testing.T) {
	// Create a buffer to capture the output
	var buf bytes.Buffer

	// Run the printToConsole function with the buffer as the output writer
	printToConsole(&buf, "Lets gooo...")

	// Check the printed output
	expected := "Lets gooo...\n"
	actual := buf.String()

	if actual != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, actual)
	}
}
