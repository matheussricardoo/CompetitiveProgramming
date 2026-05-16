// Problem: A. Marisa Steals Reimu's Takeout
// Link: https://codeforces.com/contest/2228/problem/A

package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scan(&t)
	for i := 1; i <= t; i++ {
		var n int
		fmt.Scan(&n)

		countZ := 0
		countO := 0
		countT := 0

		for j := 1; j <= n; j++ {
			var w int
			fmt.Scan(&w)
			if w == 0 {
				countZ++
			} else if w == 1 {
				countO++
			} else if w == 2 {
				countT++
			}
		}
		operations := countZ

		pairs := countO
		
		if countT < pairs {
			pairs = countT
		}
		operations += pairs

		countO -= pairs
		countT -= pairs

		operations += countO / 3
		operations += countT / 3

		fmt.Println(operations)
	}
}
