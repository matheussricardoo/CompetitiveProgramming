package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	home := make([]int, n)
	away := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&home[i], &away[i])
	}

	count := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i != j && home[i] == away[j] {
				count++
			}
		}
	}
	fmt.Println(count)
}
