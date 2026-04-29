package main

import (
	"fmt"
)

func main() {
	var a, b, count int
	fmt.Scan(&a, &b)

	for i := a; i <= b; i *= 3 {
		b = b * 2
		count++
	}
	fmt.Println(count)
}
