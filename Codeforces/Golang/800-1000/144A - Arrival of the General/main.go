package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	maxN, minN := 0, 101
	maxI, minI := 0, 0

	for i := 0; i < n; i++ {
		var heighs int
		fmt.Scan(&heighs)

		if heighs > maxN {
			maxN = heighs
			maxI = i
		}

		if heighs <= minN {
			minN = heighs
			minI = i
		}
	}

	ans := maxI + (n - 1 - minI)
	if maxI > minI {
		ans--
	}

	fmt.Println(ans)
}
