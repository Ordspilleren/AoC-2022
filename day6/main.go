package main

import (
	"log"
	"os"
)

func FindMarker(dataString string, length int) (string, int) {
	arrayLength := length - 1

	for index := range dataString {
		if index < arrayLength {
			continue
		}

		currentFour := dataString[index-arrayLength : index+1]

		runeMap := make(map[rune]bool)
		for _, subRune := range currentFour {
			runeMap[subRune] = true
		}

		if len(runeMap) == length {
			return currentFour, index + 1
		}
	}
	return "", 0
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	dataString := string(data)

	marker, index := FindMarker(dataString, 14)
	log.Printf("Marker found at index %d\nMarker is: %s", index, marker)
}
