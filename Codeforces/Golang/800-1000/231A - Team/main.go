package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	s := 0

	for i := 0; i < n; i++ {
		if n >= 1 && n <= 1000 {
			var p, v, t int
			fmt.Scan(&p, &v, &t)
			if p+v+t >= 2 {
				s++
			}
		}
	}

	fmt.Println(s)
}
