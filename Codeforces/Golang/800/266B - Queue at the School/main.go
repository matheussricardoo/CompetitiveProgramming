package main

import (
	"fmt"
)

func main() {
	var n, t int
	var s string
	fmt.Scan(&n, &t, &s)

	queue := []byte(s)

	for time := 0; time < t; time++ {
		for i := 0; i < n-1; i++ {
			if queue[i] == 'B' && queue[i+1] == 'G' {
				queue[i], queue[i+1] = 'G', 'B'
				i++
			}
		}
	}
	fmt.Println(string(queue))
}
