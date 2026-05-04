package main

import (
	"fmt"
)

func main() {
	var n string
	fmt.Scan(&n)

	luckyCount := 0
	for _, ch := range n {
		if ch == '4' || ch == '7' {
			luckyCount++
		}
	}

	if luckyCount == 4 || luckyCount == 7 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
