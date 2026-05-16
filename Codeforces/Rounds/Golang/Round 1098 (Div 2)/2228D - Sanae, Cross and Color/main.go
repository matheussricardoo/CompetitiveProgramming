// Problem: D. Sanae, Cross and Color
// Link: https://codeforces.com/contest/2228/problem/D


package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x, y int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	fmt.Fscan(reader, &t)

	for tc := 0; tc < t; tc++ {
		var n int
		fmt.Fscan(reader, &n)

		points := make([]Point, n)

		hasX := make([]bool, n+1)
		hasY := make([]bool, n+1)

		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &points[i].x, &points[i].y)
			hasX[points[i].x] = true
			hasY[points[i].y] = true
		}
		uniqueX := make([]int, 0, n)
		uniqueY := make([]int, 0, n)
		for i := 1; i <= n; i++ {
			if hasX[i] {
				uniqueX = append(uniqueX, i)
			}
			if hasY[i] {
				uniqueY = append(uniqueY, i)
			}
		}

		if len(uniqueX) < 2 || len(uniqueY) < 2 {
			fmt.Fprintln(writer, 0)
			continue
		}

		m := len(uniqueX)
		minYAtX := make([]int, m)
		maxYAtX := make([]int, m)
		for i := 0; i < m; i++ {
			minYAtX[i] = 2e9
			maxYAtX[i] = -2e9
		}
		xToIndex := make([]int, n+1)
		for i, x := range uniqueX {
			xToIndex[x] = i
		}

		for _, p := range points {
			idx := xToIndex[p.x]
			if p.y < minYAtX[idx] {
				minYAtX[idx] = p.y
			}
			if p.y > maxYAtX[idx] {
				maxYAtX[idx] = p.y
			}
		}
		prefMin := make([]int, m)
		prefMax := make([]int, m)
		prefMin[0] = minYAtX[0]
		prefMax[0] = maxYAtX[0]
		for i := 1; i < m; i++ {
			prefMin[i] = prefMin[i-1]
			if minYAtX[i] < prefMin[i] {
				prefMin[i] = minYAtX[i]
			}
			prefMax[i] = prefMax[i-1]
			if maxYAtX[i] > prefMax[i] {
				prefMax[i] = maxYAtX[i]
			}
		}
		suffMin := make([]int, m)
		suffMax := make([]int, m)
		suffMin[m-1] = minYAtX[m-1]
		suffMax[m-1] = maxYAtX[m-1]
		for i := m - 2; i >= 0; i-- {
			suffMin[i] = suffMin[i+1]
			if minYAtX[i] < suffMin[i] {
				suffMin[i] = minYAtX[i]
			}
			suffMax[i] = suffMax[i+1]
			if maxYAtX[i] > suffMax[i] {
				suffMax[i] = maxYAtX[i]
			}
		}
		getIntervalIdx := func(val int) int {
			low, high := 0, len(uniqueY)-1
			ans := -1
			for low <= high {
				mid := low + (high-low)/2
				if uniqueY[mid] <= val {
					ans = mid
					low = mid + 1
				} else {
					high = mid - 1
				}
			}
			return ans
		}

		var totalDistinctColorings int64 = 0

		for i := 0; i < m-1; i++ {
			minL, maxL := prefMin[i], prefMax[i]
			minR, maxR := suffMin[i+1], suffMax[i+1]

			lowY := minL
			if minR > lowY {
				lowY = minR
			}

			highY := maxL
			if maxR < highY {
				highY = maxR
			}
			highY--

			if lowY <= highY {
				lowIdx := getIntervalIdx(lowY)
				highIdx := getIntervalIdx(highY)
				if highIdx >= lowIdx {
					totalDistinctColorings += int64(highIdx - lowIdx + 1)
				}
			}
		}

		fmt.Fprintln(writer, totalDistinctColorings)
	}
}
