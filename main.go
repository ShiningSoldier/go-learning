package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

func main() {
	var encryptedText = "Olssv, tf mypluk. Ovd hyl fvb kvpun?"
	var decryptedText = []string{"Doing", "how", "you", "friend", "my"}
	var decryptedString = strings.Join(decryptedText, " ")
	decryptedString = strings.ToLower(decryptedString)
	decryptedText = strings.Split(decryptedString, " ")

	fmt.Printf(decryptString(encryptedText, decryptedText))
}

func decryptString(encryptedText string, decryptedText []string) string {
	s := strings.ToLower(encryptedText)
	sort.Strings(decryptedText)

	var step byte = 0
	result := decrypt(s, step)

	for compareTexts(result, decryptedText) != true {
		step++
		result = decrypt(s, step)
	}

	return fmt.Sprintf("Full text: %s, key: %d", result, step)
}

func compareTexts(result string, decryptedText []string) bool {
	splitResult := strings.Split(result, " ")
	firstWord := decryptedText[0]
	secondWord := decryptedText[1]
	thirdWord := decryptedText[2]

	if contains(splitResult, firstWord) == false && contains(splitResult, secondWord) == false && contains(splitResult, thirdWord) == false {
		return false
	}

	return true
}

func decrypt(s string, step byte) string {
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

func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}
