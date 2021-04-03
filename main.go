package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var toCompress string = "Helllllo world"
	var toDecompress string = "He#5#lo world"

	fmt.Printf(compress(toCompress))
	fmt.Printf(decompress(toDecompress))

}

func compress(s string) string {
	var strLen = len(s)
	var result []string
	var i = 0

	for i = 0; i < strLen; i++ {
		var tempCount = 1
		for j := i + 1; j < strLen; j++ {
			if s[i] != s[j] {
				break
			}
			tempCount++
		}

		if tempCount > 4 {
			var mask = fmt.Sprintf("#%d#%s", tempCount, string(s[i]))
			result = append(result, mask)
			i = i + tempCount - 1
		} else {
			result = append(result, string(s[i]))
		}
	}

	return strings.Join(result, "")
}

func decompress(s string) string {
	var regString = regexp.MustCompile(`#\d#[a-zA-Z]`)
	var foundString = regString.FindString(s)
	var regNum = regexp.MustCompile(`(\d)`)
	var foundNum = regNum.FindString(foundString)
	var lastCharacter = foundString[len(foundString)-1:]
	strToNum, err := strconv.Atoi(foundNum)
	if err != nil {
		log.Fatal(err)
	}
	tempArray := make([]string, strToNum)

	for i := range tempArray {
		tempArray[i] = lastCharacter
	}

	var result = regString.ReplaceAllString(s, strings.Join(tempArray, ""))

	return result
}
