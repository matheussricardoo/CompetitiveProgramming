package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var n string
		fmt.Scan(&n)

		sumFirsts := int(n[0]) + int(n[1]) + int(n[2])
		sumLasts := int(n[3]) + int(n[4]) + int(n[5])

		if sumFirsts == sumLasts {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}

}
