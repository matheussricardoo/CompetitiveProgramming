package main

import (
	"fmt"
)

func main() {
	var s1, s2, pile string
	fmt.Scan(&s1, &s2, &pile)

	counts := make(map[rune]int)

	for _, char := range s1 + s2 {
		counts[char]++
	}

	for _, char := range pile {
		counts[char]--
	}

	for _, count := range counts {
		if count != 0 {
			fmt.Println("NO")
			return
		}
	}

	if len(s1)+len(s2) != len(pile) {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}
}
