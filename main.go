package main

import (
	"fmt"
	"log"
	"regexp"
	"sort"
	"strings"
)

func main() {
	var s = "один, два - это 2, три один два, много слов: один"

	fmt.Print(countAndSort(s))
}

func countAndSort(s string) string {
	reg, err := regexp.Compile("[^а-яА-Я0-9 ]+")
	var tempMap = make(map[string]int)
	var result string

	if err != nil {
		log.Fatal(err)
	}
	var processedString = reg.ReplaceAllString(s, "")

	regSpaces, err := regexp.Compile(`\s+`)
	if err != nil {
		log.Fatal(err)
	}
	processedString = regSpaces.ReplaceAllString(processedString, " ")

	var stringToSlice = strings.Split(processedString, " ")

	for _, v := range stringToSlice {
		_, wordExist := tempMap[v]
		if wordExist {
			tempMap[v]++
		} else {
			tempMap[v] = 1
		}
	}

	var sliceForSorting = make([]string, 0, len(tempMap))
	for i := range tempMap {
		sliceForSorting = append(sliceForSorting, string(i))
	}

	sort.Slice(sliceForSorting, func(i, j int) bool {
		return tempMap[sliceForSorting[i]] > tempMap[sliceForSorting[j]]
	})

	for _, k := range sliceForSorting {
		result += fmt.Sprintf("%s(%d) ", k, tempMap[k])
	}

	return result
}
