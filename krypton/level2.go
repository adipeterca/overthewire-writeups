package main

import (
	"fmt"
	"strings"
)

/*
 */
func level2() {

	encryptedString := "OMQEMDUEQMEK"

	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// during encryption, A -> M
	// no difference in uppercase and lowercase

	difference := int('M' - 'A')

	decodedBytes := make([]byte, len(encryptedString))

	for index, value := range encryptedString {

		newValue := byte(' ')
		if value != ' ' {
			alpIndex := strings.Index(alphabet, string(value))
			newIndex := (alpIndex - difference) % len(alphabet)
			newIndex = (newIndex + len(alphabet)) % len(alphabet)

			newValue = alphabet[newIndex]
		}
		decodedBytes[index] = newValue

	}

	decodedString := string(decodedBytes)

	fmt.Println(decodedString)
}
