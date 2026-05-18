package main

import (
	"fmt"
)

func main() {
	var n, count, polices int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var n int
		fmt.Scan(&n)
		if n > 0 {
			polices += n
		} else {
			if polices > 0 {
				polices--
			} else {
				count++
			}
		}
	}
	fmt.Println(count)
}
