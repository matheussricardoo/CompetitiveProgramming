package main

import (
	"fmt"
	"slices"
)

func main() {
	var s, t string
	fmt.Scan(&s, &t)
	inverted := []rune(s)
	slices.Reverse(inverted)

	if string(inverted) == t {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
