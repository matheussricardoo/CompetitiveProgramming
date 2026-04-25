package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n)
	fmt.Scan(&k)

	a := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	targetscore := a[k-1]
	count := 0

	for i := 0; i < n; i++ {
		if a[i] >= targetscore && a[i] > 0 {
			count++
		} else {
			break
		}
	}
	fmt.Println(count)
}
