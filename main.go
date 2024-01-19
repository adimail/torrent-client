// main.go
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	outputWriter := os.Stdout
	printToConsole(outputWriter, "Lets gooo...")

	// fmt.Println(decodeBencode("l3:one3:twoe"))
}

func printToConsole(w io.Writer, message string) {
	fmt.Fprintln(w, message)
}
