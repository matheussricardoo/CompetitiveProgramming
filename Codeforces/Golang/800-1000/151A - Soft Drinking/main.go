package main

import (
	"fmt"
)

func main() {
	var n, k, l, c, d, p, nl, np int
	fmt.Scan(&n, &k, &l, &c, &d, &p, &nl, &np)
	mlDrink := k * l
	tostas := mlDrink / nl
	lemons := c * d
	salt := p / np
	takeItLight := min(tostas, lemons, salt) / n
	fmt.Println(takeItLight)
}
