package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)
	parts := strings.Split(s, "+")
	slices.Sort(parts)
	if len(parts) == 1 {
		fmt.Println(parts[0])
	} else {
		fmt.Println(strings.Join(parts, "+"))
	}

}
