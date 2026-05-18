package main

import (
	"fmt"
)

func main() {
	var y int
	fmt.Scan(&y)

	for {
		y++

		a := y / 1000
		b := (y / 100) % 10
		c := (y / 10) % 10
		d := y % 10

		if a != b && a != c && a != d && b != c && b != d && c != d {
			fmt.Println(y)
			break
		}
	}

}
