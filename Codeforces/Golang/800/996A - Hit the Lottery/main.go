package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	bills := []int{100, 20, 10, 5, 1}
	count := 0

	for _, bill := range bills {
		count += n / bill
		n %= bill
	}

	fmt.Println(count)
}
