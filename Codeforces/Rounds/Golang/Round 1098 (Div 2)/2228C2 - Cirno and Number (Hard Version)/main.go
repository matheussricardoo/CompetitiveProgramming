// Problem: C2. Cirno and Number (Hard Version)
// Link: https://codeforces.com/contest/2228/problem/C2

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

		d := make([]int64, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&d[j])
		}

		a, _ := strconv.ParseInt(aStr, 10, 64)
		length := len(aStr)

		candidates := make(map[int64]bool)

		if n == 1 && d[0] == 0 {
			candidates[0] = true
		} else {
			if length > 1 {
				var num int64 = 0
				for j := 0; j < length-1; j++ {
					num = num*10 + d[n-1]
				}
				candidates[num] = true
			} else {
				if d[0] == 0 {
					candidates[0] = true
				}
			}
			{
				var num int64 = 0
				if d[0] > 0 {
					for j := 0; j < length+1; j++ {
						num = num*10 + d[0]
					}
				} else if n > 1 {
					num = d[1]
					for j := 0; j < length; j++ {
						num = num*10 + d[0]
					}
				}
				if num > 0 {
					candidates[num] = true
				}
			}
			var prefix int64 = 0
			isExactMatchPossible := true

			for j := 0; j < length; j++ {
				targetDigit := int64(aStr[j] - '0')
				for k := 0; k < n; k++ {
					dk := d[k]

					if j == 0 && dk == 0 && length > 1 {
						continue
					}

					num := prefix*10 + dk
					if dk < targetDigit {
						for m := j + 1; m < length; m++ {
							num = num*10 + d[n-1]
						}
						candidates[num] = true
					} else if dk > targetDigit {
						for m := j + 1; m < length; m++ {
							num = num*10 + d[0]
						}
						candidates[num] = true
					}
				}
				found := false
				for k := 0; k < n; k++ {
					if d[k] == targetDigit {
						found = true
						break
					}
				}

				if found {
					prefix = prefix*10 + targetDigit
				} else {
					isExactMatchPossible = false
					break
				}
			}

			if isExactMatchPossible {
				candidates[a] = true
			}
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
