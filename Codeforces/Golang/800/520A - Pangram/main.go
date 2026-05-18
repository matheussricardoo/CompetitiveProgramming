package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	var s string
	fmt.Scan(&n, &s)

	if n < 26 {
		fmt.Println("NO")
		return
	}

	letters := make(map[rune]bool)

	s = strings.ToLower(s)

	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			letters[r] = true
		}
	}

	if len(letters) == 26 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
