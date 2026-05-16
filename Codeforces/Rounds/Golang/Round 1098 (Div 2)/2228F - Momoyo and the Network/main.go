// Problem: F. Momoyo and the Network
// Link: https://codeforces.com/contest/2228/problem/F

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	n, k int

	adj            [][]int
	children       [][]int
	sortedChildren [][]int

	parent []int
	order  []int

	sub   []int64
	total int64
)

func check(X int64) bool {
	f := make([]int, n+1)

	for idx := n - 1; idx >= 0; idx-- {
		u := order[idx]
		best := -1

		if sub[u] >= X {
			best = 0
		}

		for _, v := range children[u] {
			if sub[u]-sub[v] >= X && f[v] >= 0 {
				if cand := 1 + f[v]; cand > best {
					best = cand
				}
			}
		}

		f[u] = best
	}
	for u := 1; u <= n; u++ {
		ch := sortedChildren[u]
		d := len(ch)
		if d == 0 {
			continue
		}
		for _, v := range ch {
			if f[v] < 0 {
				continue
			}
			capLen := 1 + f[v]
			if capLen >= k && total-sub[v] >= X {
				return true
			}
		}
		if k < 2 || d < 2 {
			continue
		}

		caps := make([]int, d)
		for i, v := range ch {
			if f[v] >= 0 {
				caps[i] = 1 + f[v]
			} else {
				caps[i] = -1
			}
		}

		size := 1
		for size < d {
			size <<= 1
		}
		seg := make([]int, 2*size)
		for i := range seg {
			seg[i] = -1
		}
		for i := 0; i < d; i++ {
			seg[size+i] = caps[i]
		}
		for i := size - 1; i >= 1; i-- {
			if seg[i<<1] > seg[i<<1|1] {
				seg[i] = seg[i<<1]
			} else {
				seg[i] = seg[i<<1|1]
			}
		}

		queryMax := func(l, r int) int {
			if l > r {
				return -1
			}
			l += size
			r += size
			res := -1
			for l <= r {
				if l&1 == 1 {
					if seg[l] > res {
						res = seg[l]
					}
					l++
				}
				if r&1 == 0 {
					if seg[r] > res {
						res = seg[r]
					}
					r--
				}
				l >>= 1
				r >>= 1
			}
			return res
		}

		limit := total - X
		r := d - 1

		for i := 0; i < d-1; i++ {
			for r > i && sub[ch[i]]+sub[ch[r]] > limit {
				r--
			}
			if r <= i {
				break
			}
			if caps[i] < 1 {
				continue
			}
			bestOther := queryMax(i+1, r)
			if bestOther >= 1 && caps[i]+bestOther >= k {
				return true
			}
		}
	}

	return false
}

func main() {
	in := bufio.NewReaderSize(os.Stdin, 1<<20)
	out := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	for ; t > 0; t-- {
		fmt.Fscan(in, &n, &k)

		adj = make([][]int, n+1)
		children = make([][]int, n+1)
		sortedChildren = make([][]int, n+1)
		parent = make([]int, n+1)
		sub = make([]int64, n+1)
		order = make([]int, 0, n)
		total = 0

		weights := make([]int64, n+1)
		for i := 1; i <= n; i++ {
			fmt.Fscan(in, &weights[i])
			total += weights[i]
		}

		for i := 0; i < n-1; i++ {
			var u, v int
			fmt.Fscan(in, &u, &v)
			adj[u] = append(adj[u], v)
			adj[v] = append(adj[v], u)
		}
		stack := []int{1}
		parent[1] = -1
		for len(stack) > 0 {
			u := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			order = append(order, u)

			for _, v := range adj[u] {
				if v == parent[u] {
					continue
				}
				parent[v] = u
				children[u] = append(children[u], v)
				stack = append(stack, v)
			}
		}
		for i := n - 1; i >= 0; i-- {
			u := order[i]
			sub[u] = weights[u]
			for _, v := range children[u] {
				sub[u] += sub[v]
			}
		}
		for u := 1; u <= n; u++ {
			if len(children[u]) == 0 {
				continue
			}
			arr := append([]int(nil), children[u]...)
			sort.Slice(arr, func(i, j int) bool {
				return sub[arr[i]] < sub[arr[j]]
			})
			sortedChildren[u] = arr
		}

		var lo int64 = 1
		var hi int64 = total
		var ans int64 = 0

		for lo <= hi {
			mid := lo + (hi-lo)/2
			if check(mid) {
				ans = mid
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}

		if ans == 0 {
			fmt.Fprintln(out, -1)
		} else {
			fmt.Fprintln(out, ans)
		}
	}
}
