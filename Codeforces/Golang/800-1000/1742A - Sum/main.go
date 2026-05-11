package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var a, b, c int
		fmt.Scan(&a, &b, &c)
		if a+b == c || a+c == b || c+b == a {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
