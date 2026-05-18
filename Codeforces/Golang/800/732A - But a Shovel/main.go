package main

import (
	"fmt"
)

func main() {
	var k, r int
	fmt.Scan(&k, &r)

	for shovels := 1; shovels <= 10; shovels++ {
		totalPrice := shovels * k
		if totalPrice%10 == 0 || totalPrice%10 == r {
			fmt.Println(shovels)
			break
		}
	}
}
