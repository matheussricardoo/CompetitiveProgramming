package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	walk := x / 5
	if x%5 != 0 {
		walk++
	}

	fmt.Println(walk)

}
