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

		if n%3 != 0 {
			fmt.Println("First")
		} else {
			fmt.Println("Second")
		}

	}
}
