package main

import (
	"fmt"
)

func main() {
	var players string
	fmt.Scan(&players)

	if len(players) < 7 {
		fmt.Println("NO")
		return
	}
	count := 1
	dangerous := false
	for i := 1; i < len(players); i++ {
		if players[i] == players[i-1] {
			count++
			if count == 7 {
				dangerous = true
				break
			}
		} else {
			count = 1
		}
	}

	if dangerous {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
