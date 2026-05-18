package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	levels := make(map[int]struct{})

	var p int
	fmt.Scan(&p)

	for i := 0; i < p; i++ {
		var l int
		fmt.Scan(&l)
		levels[l] = struct{}{}
	}

	var q int
	fmt.Scan(&q)
	for i := 0; i < q; i++ {
		var l int
		fmt.Scan(&l)
		levels[l] = struct{}{}
	}

	if len(levels) == n {
		fmt.Println("I become the guy.")
	} else {
		fmt.Println("Oh, my keyboard!")
	}
}
