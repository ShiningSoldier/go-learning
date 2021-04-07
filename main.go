package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	var s = "J'n mfbsojoh hp qsphsbnnjoh mbohvbhf"
	fmt.Print(decrypt(s, 1))
}

func decrypt(s string, step byte) string {
	s = strings.ToLower(s)
	strToBytes := []byte(s)

	var result = make([]byte, len(strToBytes), len(strToBytes))
	var isLetter = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

	for i := range strToBytes {
		if isLetter(string(strToBytes[i])) {
			if strToBytes[i] >= 97+step {
				result[i] = strToBytes[i] - step
			} else {
				result[i] = strToBytes[i] + 26 - step
			}
		} else {
			result[i] = strToBytes[i]
		}
	}

	return string(result)
}
