package main

import (
	"fmt"
)

func main() {
	var n, sum int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var s string
		fmt.Scan(&s)
		switch s {
		case "Tetrahedron":
			sum += 4
		case "Cube":
			sum += 6
		case "Octahedron":
			sum += 8
		case "Dodecahedron":
			sum += 12
		case "Icosahedron":
			sum += 20
		}
	}
	fmt.Println(sum)
}
