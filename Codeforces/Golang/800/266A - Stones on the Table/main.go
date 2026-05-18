package main

import (
	"fmt"
)

func main() {
	var n int
	var s string
	fmt.Scan(&n, &s)

	remove := 0
	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			remove++
		}
	}
	fmt.Println(remove)
}
