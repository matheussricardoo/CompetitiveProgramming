package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var input string
	count := 0

	for i := 0; i < n; i++ {
		fmt.Scan(&input)
		if input == "++X" || input == "X++" {
			count += 1
		} else {
			count -= 1
		}
	}
	fmt.Println(count)
}
