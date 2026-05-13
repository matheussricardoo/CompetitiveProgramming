package main

import (
	"fmt"
)

func main() {
	var n, count int
	fmt.Scan(&n)

	var score int
	fmt.Scan(&score)

	max := score
	min := score

	for i := 1; i < n; i++ {
		fmt.Scan(&score)
		if score > max {
			max = score
			count++
		} else if score < min {
			min = score
			count++
		}
	}
	fmt.Println(count)
}
