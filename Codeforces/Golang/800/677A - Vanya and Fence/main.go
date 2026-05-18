package main

import (
	"fmt"
)

func main() {
	var n, h, l int
	fmt.Scan(&n, &h)
	for i := 0; i < n; i++ {
		var j int
		fmt.Scan(&j)

		if j <= h {
			l++
		} else {
			l += 2
		}
	}
	fmt.Println(l)
}
