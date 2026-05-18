package main

import (
	"fmt"
	"math"
)

func main() {
	var v int
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			fmt.Scan(&v)
			if v == 1 {
				row := math.Abs(float64(i - 3))
				column := math.Abs(float64(j - 3))
				result := row + column
				fmt.Println(int(result))
				return
			}
		}
	}
}
