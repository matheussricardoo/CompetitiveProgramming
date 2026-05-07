package main

import (
	"fmt"
	"math"
)

func main() {
	var n, p, q, count, notNegative int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&p, &q)
		sizeRoom := p - q
		notNegative = int(math.Abs(float64(sizeRoom)))
		if notNegative >= 2 {
			count++
		}
	}
	fmt.Println(count)
}
