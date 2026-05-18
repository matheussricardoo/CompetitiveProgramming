package main

import (
	"fmt"
)

func main() {
	var t, rating int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		fmt.Scan(&rating)
		if 1900 <= rating {
			fmt.Println("Division 1")
		} else if 1600 <= rating && rating <= 1899 {
			fmt.Println("Division 2")
		} else if 1400 <= rating && rating <= 1599 {
			fmt.Println("Division 3")
		} else if rating <= 1399 {
			fmt.Println("Division 4")
		}
	}
}
