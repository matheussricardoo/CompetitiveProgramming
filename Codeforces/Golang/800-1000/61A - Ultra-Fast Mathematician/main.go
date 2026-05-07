package main

import (
	"fmt"
)

func main() {
	var n1, n2 string
	fmt.Scan(&n1, &n2)

	result := make([]byte, len(n1))

	for i := 0; i < len(n1); i++ {
		if n1[i] != n2[i] {
			result[i] = '1'
		} else {
			result[i] = '0'
		}
	}
	fmt.Println(string(result))
}
