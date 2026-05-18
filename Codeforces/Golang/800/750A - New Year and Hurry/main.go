package main

import (
	"fmt"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	timeRemaining := 240 - k
	solvedProblems := 0

	for i := 1; i <= n; i++ {
		timeToProblem := 5 * i

		if timeRemaining >= timeToProblem {
			timeRemaining -= timeToProblem
			solvedProblems++
		} else {
			break
		}
	}
	fmt.Println(solvedProblems)
}
