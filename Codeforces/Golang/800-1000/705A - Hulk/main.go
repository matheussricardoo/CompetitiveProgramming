package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	var sb strings.Builder

	for i := 1; i <= n; i++ {
		if i%2 != 0 {
			sb.WriteString("I hate")
		} else {
			sb.WriteString("I love")
		}

		if i == n {
			sb.WriteString(" it")
		} else {
			sb.WriteString(" that ")
		}
	}
	fmt.Println(sb.String())
}
