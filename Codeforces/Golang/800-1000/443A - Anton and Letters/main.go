package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	set := make(map[rune]struct{})

	for _, char := range input {
		if char >= 'a' && char <= 'z' {
			set[char] = struct{}{}
		}
	}

	fmt.Println(len(set))

}
