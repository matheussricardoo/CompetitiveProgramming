// Problem: C1. Cirno and Number (Easy Version)
// Link: https://codeforces.com/contest/2228/problem/C1

package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var aStr string
		var n int
		fmt.Scan(&aStr, &n)

		var d1, d2 int64
		fmt.Scan(&d1, &d2)

		a, _ := strconv.ParseInt(aStr, 10, 64)
		length := len(aStr)

		candidates := make(map[int64]bool)

		if length > 1 {
			var num int64 = 0
			for j := 0; j < length-1; j++ {
				num = num*10 + d2
			}
			candidates[num] = true
		} else {
			if d1 == 0 {
				candidates[0] = true
			}
		}
		var num int64 = 0
		if d1 > 0 {
			for j := 0; j < length+1; j++ {
				num = num*10 + d1
			}
		} else {
			num = d2
			for j := 0; j < length; j++ {
				num = num * 10
			}
		}
		candidates[num] = true
		var prefix int64 = 0
		isExactMatchPossible := true

		for j := 0; j < length; j++ {
			targetDigit := int64(aStr[j] - '0')

			if j > 0 || d1 > 0 || length == 1 {
				num := prefix*10 + d1
				if d1 < targetDigit {
					for k := j + 1; k < length; k++ {
						num = num*10 + d2
					}
					candidates[num] = true
				} else if d1 > targetDigit {
					for k := j + 1; k < length; k++ {
						num = num*10 + d1
					}
					candidates[num] = true
				}
			}
			{
				num := prefix*10 + d2
				if d2 < targetDigit {
					for k := j + 1; k < length; k++ {
						num = num*10 + d2
					}
					candidates[num] = true
				} else if d2 > targetDigit {
					for k := j + 1; k < length; k++ {
						num = num*10 + d1
					}
					candidates[num] = true
				}
			}
			if targetDigit == d1 && (j > 0 || d1 > 0 || length == 1) {
				prefix = prefix*10 + d1
			} else if targetDigit == d2 {
				prefix = prefix*10 + d2
			} else {
				isExactMatchPossible = false
				break
			}
		}
		if isExactMatchPossible {
			candidates[a] = true
		}
		var minDiff int64 = math.MaxInt64
		for c := range candidates {
			diff := a - c
			if diff < 0 {
				diff = -diff
			}
			if diff < minDiff {
				minDiff = diff
			}
		}
		fmt.Println(minDiff)
	}
}
