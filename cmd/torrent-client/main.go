package main

import (
	"fmt"

	"github.com/adimail/torrent-client/internal/bencode"
)

func main() {
	bencodedData := "l4:spam4:eggse"

	decoder := bencode.DecodeBencode
	result, err := decoder(bencodedData)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Decoded Result:", result)
	}
}
