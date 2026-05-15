package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	cards := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&cards[i])
	}

	left, right := 0, n-1
	serejaScore, dimaScore := 0, 0
	isSerejaTurn := true

	for left <= right {
		var pickedCard int
		if cards[left] > cards[right] {
			pickedCard = cards[left]
			left++
		} else {
			pickedCard = cards[right]
			right--
		}

		if isSerejaTurn {
			serejaScore += pickedCard
		} else {
			dimaScore += pickedCard
		}
		isSerejaTurn = !isSerejaTurn
	}

	fmt.Printf("%d %d\n", serejaScore, dimaScore)
}
