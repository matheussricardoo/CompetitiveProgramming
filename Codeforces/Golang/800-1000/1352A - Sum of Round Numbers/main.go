package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var n int
		fmt.Scan(&n)

		var results []int
		multiplier := 1

		for n > 0 {
			digit := n % 10
			if digit != 0 {
				results = append(results, digit*multiplier)
			}
			n /= 10
			multiplier *= 10
		}
		fmt.Println(len(results))
		for j, val := range results {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(val)
		}
		fmt.Println()
	}
}
