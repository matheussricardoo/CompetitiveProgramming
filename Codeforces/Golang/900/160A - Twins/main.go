package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	if !scanner.Scan() {
		return
	}
	n, _ := strconv.Atoi(scanner.Text())

	coins := make([]int, n)
	totalSum := 0

	for i := 0; i < n; i++ {
		scanner.Scan()
		coins[i], _ = strconv.Atoi(scanner.Text())
		totalSum += coins[i]
	}
	sort.Ints(coins)

	mySum := 0
	count := 0

	for i := n - 1; i >= 0; i-- {
		mySum += coins[i]
		count++

		if mySum > totalSum/2 {
			break
		}
	}
	fmt.Println(count)
}
