package main

import (
	"fmt"
)

func main() {
	var n, juice, sum float64
	fmt.Scan(&n)
	for i := 0; i < int(n); i++ {
		fmt.Scan(&juice)
		sum += juice
	}
	sum = sum / n
	fmt.Println(sum)
}
