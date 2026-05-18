package main

import "fmt"

func main() {
	var repeat int
	var word string
	fmt.Scan(&repeat)

	for i := 0; i < repeat; i++ {
		fmt.Scan(&word)
		size := len(word)
		if size > 10 {
			fmt.Printf("%s%d%s\n", word[0:1], size-2, word[size-1:])
		} else {
			fmt.Println(word)
		}
	}
}
