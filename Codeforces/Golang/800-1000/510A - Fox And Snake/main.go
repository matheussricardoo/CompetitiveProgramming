package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	right := true

	for i := 1; i <= n; i++ {
		if i%2 != 0 {
			for j := 0; j < m; j++ {
				fmt.Print("#")
			}
		} else {
			if right {
				for j := 0; j < m-1; j++ {
					fmt.Print(".")
				}
				fmt.Print("#")
			} else {
				fmt.Print("#")
				for j := 0; j < m-1; j++ {
					fmt.Print(".")
				}
			}
			right = !right
		}
		fmt.Println()
	}
}
