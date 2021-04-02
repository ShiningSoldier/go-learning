package main

import "fmt"

func main() {
	fmt.Println(isNumberAutomorphic(2))
}

func isNumberAutomorphic(num int) bool {
	var numPowered = num * num

	for num > 0 {
		if num%10 != numPowered%10 {
			return false
		}

		num /= 10
		numPowered /= 10
	}

	return true
}
