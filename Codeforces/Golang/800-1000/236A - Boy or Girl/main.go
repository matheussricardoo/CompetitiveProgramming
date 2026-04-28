package main

import "fmt"

func main() {
	var name string
	fmt.Scan(&name)

	distinctCh := make(map[rune]bool)

	for _, char := range name {
		distinctCh[char] = true
	}
	count := len(distinctCh)

	if count%2 == 0 {
		fmt.Println("CHAT WITH HER!")
	} else {
		fmt.Println("IGNORE HIM!")
	}
}
