package main

import (
	"fmt"
	"strings"
)

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var s string
		fmt.Scan(&s)

		if strings.ToLower(s) == "yes" {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
