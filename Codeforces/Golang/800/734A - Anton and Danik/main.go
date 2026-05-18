package main

import (
	"fmt"
)

func main() {
	var n, a, d int
	var s string
	fmt.Scan(&n, &s)

	for _, ch := range s {
		if ch == 'A' {
			a++
		} else {
			d++
		}
	}

	if a == d {
		fmt.Println("Friendship")
	} else if a > d {
		fmt.Println("Anton")
	} else {
		fmt.Println("Danik")
	}
}
