package main

import (
	"fmt"
	"sort"
)

func main() {
	points := make([]int, 3)
	fmt.Scan(&points[0], &points[1], &points[2])

	sort.Ints(points)

	distance := points[2] - points[0]

	fmt.Println(distance)

}
