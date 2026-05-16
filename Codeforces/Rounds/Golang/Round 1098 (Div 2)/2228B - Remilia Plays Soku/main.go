// Problem: B. Remilia Plays Soku
// Link: https://codeforces.com/contest/2228/problem/B

package main

import (
	"fmt"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var n, x1, x2, k int
		fmt.Scan(&n, &x1, &x2, &k)

		directDist := abs(x1 - x2)
		shortestDist := min(directDist, n-directDist)

		var totalSeconds int
		if n <= 3 {
			totalSeconds = shortestDist
		} else {
			totalSeconds = shortestDist + k
		}
		fmt.Println(totalSeconds)
	}
}
