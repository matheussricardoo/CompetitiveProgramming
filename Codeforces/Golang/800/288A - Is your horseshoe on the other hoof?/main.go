package main

import (
	"fmt"
)

func main() {
	var s1, s2, s3, s4 int
	fmt.Scan(&s1, &s2, &s3, &s4)

	set := make(map[int]struct{})

	set[s1] = struct{}{}
	set[s2] = struct{}{}
	set[s3] = struct{}{}
	set[s4] = struct{}{}

	distinct := len(set)
	result := 4 - distinct

	fmt.Println(result)

}
