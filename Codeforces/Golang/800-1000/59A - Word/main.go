package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	var s string
	fmt.Scan(&s)

	upperCount := 0
	lowerCount := 0

	for _, r := range s {
		if unicode.IsUpper(r) {
			upperCount++
		} else {
			lowerCount++
		}
	}

	if upperCount > lowerCount {
		fmt.Println(strings.ToUpper(s))
	} else {
		fmt.Println(strings.ToLower(s))
	}
}
