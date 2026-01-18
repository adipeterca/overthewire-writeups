package main

import (
	"fmt"
	"strings"
)

/**
* Sources:
* https://en.wikipedia.org/wiki/ROT13
 */
func level1() {

	encodedString := "YRIRY GJB CNFFJBEQ EBGGRA"
	decodedBytes := make([]byte, len(encodedString))

	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	for index, value := range encodedString {

		newValue := byte(' ')
		if value != ' ' {
			alpIndex := strings.Index(alphabet, string(value))
			newIndex := (alpIndex - 13) % 26
			newIndex = (newIndex + 26) % 26

			newValue = alphabet[newIndex]
		}
		decodedBytes[index] = newValue

	}

	decodedString := string(decodedBytes)

	fmt.Println(decodedString)
}
