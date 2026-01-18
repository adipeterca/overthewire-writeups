package main

import (
	"encoding/base64"
	"fmt"
)

func level0() {

	// Base64 encoding
	encodedString := "S1JZUFRPTklTR1JFQVQ="

	decodedBytes, err := base64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		panic(err)
	}

	decodedString := string(decodedBytes)

	fmt.Println("Decoded string: '" + decodedString + "'")
}
