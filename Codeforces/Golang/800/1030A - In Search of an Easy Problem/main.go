package main

import (
	"fmt"
	"slices"
)

func main() {
	var n int
	fmt.Scan(&n)

	arry := []int{}

	for i := 0; i < n; i++ {
		var j int
		fmt.Scan(&j)
		arry = append(arry, j)
	}
	oneOrZero := slices.Contains(arry, 1)

	if oneOrZero {
		fmt.Println("HARD")
	} else {
		fmt.Println("EASY")
	}
}
