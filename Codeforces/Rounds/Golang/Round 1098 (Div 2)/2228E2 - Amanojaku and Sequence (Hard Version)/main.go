// Problem: E2. Amanojaku and Sequence (Hard Version)
// Link: https://codeforces.com/contest/2228/problem/E2

package main

import (
	"bufio"
	"fmt"
	"os"
)

const MOD = 998244353

type Node struct {
	totalFixed int64
	totalFree  int64
	length     int64
	sumF       int64
	sumF2      int64
	sumC       int64
	sumC2      int64
	sumFC      int64
}

var fact []int64
var invFact []int64
var tree []Node
var baseArray []int64

func power(base, exp int64) int64 {
	res := int64(1)
	base %= MOD
	for exp > 0 {
		if exp%2 == 1 {
			res = (res * base) % MOD
		}
		base = (base * base) % MOD
		exp /= 2
	}
	return res
}

func modInverse(n int64) int64 {
	return power(n, MOD-2)
}

func precompute(maxN int) {
	fact = make([]int64, maxN+1)
	invFact = make([]int64, maxN+1)
	fact[0] = 1
	invFact[0] = 1
	for i := 1; i <= maxN; i++ {
		fact[i] = (fact[i-1] * int64(i)) % MOD
	}
	invFact[maxN] = modInverse(fact[maxN])
	for i := maxN - 1; i >= 1; i-- {
		invFact[i] = (invFact[i+1] * int64(i+1)) % MOD
	}
}

func nCr(n, r int) int64 {
	if r < 0 || r > n {
		return 0
	}
	return fact[n] * invFact[r] % MOD * invFact[n-r] % MOD
}

func createLeaf(val int64) Node {
	var node Node
	node.length = 1
	if val >= 0 {
		node.totalFixed = val
		vMod := val % MOD
		node.sumF = vMod
		node.sumF2 = (vMod * vMod) % MOD
	} else {
		node.totalFree = 1
		node.sumC = 1
		node.sumC2 = 1
	}
	return node
}

func mergeNodes(L, R Node) Node {
	var res Node
	res.totalFixed = L.totalFixed + R.totalFixed
	res.totalFree = L.totalFree + R.totalFree
	res.length = L.length + R.length

	leftFixedMod := L.totalFixed % MOD
	leftFreeMod := L.totalFree % MOD

	res.sumF = (L.sumF + R.sumF + (R.length*leftFixedMod)%MOD) % MOD

	res.sumF2 = (L.sumF2 + R.sumF2 +
		(R.length*((leftFixedMod*leftFixedMod)%MOD))%MOD +
		(2*leftFixedMod%MOD)*R.sumF%MOD) % MOD

	res.sumC = (L.sumC + R.sumC + (R.length*leftFreeMod)%MOD) % MOD

	res.sumC2 = (L.sumC2 + R.sumC2 +
		(R.length*((leftFreeMod*leftFreeMod)%MOD))%MOD +
		(2*leftFreeMod%MOD)*R.sumC%MOD) % MOD

	res.sumFC = (L.sumFC + R.sumFC +
		(((R.length*leftFixedMod)%MOD)*leftFreeMod)%MOD +
		(leftFixedMod*R.sumC)%MOD +
		(leftFreeMod*R.sumF)%MOD) % MOD

	return res
}

func buildTree(node, start, end int) {
	if start == end {
		tree[node] = createLeaf(baseArray[start])
		return
	}
	mid := (start + end) / 2
	buildTree(2*node, start, mid)
	buildTree(2*node+1, mid+1, end)
	tree[node] = mergeNodes(tree[2*node], tree[2*node+1])
}

func updateTree(node, start, end, idx int, val int64) {
	if start == end {
		baseArray[start] = val
		tree[node] = createLeaf(val)
		return
	}
	mid := (start + end) / 2
	if idx <= mid {
		updateTree(2*node, start, mid, idx, val)
	} else {
		updateTree(2*node+1, mid+1, end, idx, val)
	}
	tree[node] = mergeNodes(tree[2*node], tree[2*node+1])
}

func queryTree(node, start, end, l, r int) Node {
	if l <= start && end <= r {
		return tree[node]
	}
	mid := (start + end) / 2
	if r <= mid {
		return queryTree(2*node, start, mid, l, r)
	}
	if l > mid {
		return queryTree(2*node+1, mid+1, end, l, r)
	}
	leftNode := queryTree(2*node, start, mid, l, r)
	rightNode := queryTree(2*node+1, mid+1, end, l, r)
	return mergeNodes(leftNode, rightNode)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	precompute(1500000)

	var t int
	if _, err := fmt.Fscan(reader, &t); err != nil {
		return
	}

	for tc := 0; tc < t; tc++ {
		var n, q int
		fmt.Fscan(reader, &n, &q)

		baseArray = make([]int64, n+1)
		tree = make([]Node, 4*n+1)

		for i := 1; i <= n; i++ {
			fmt.Fscan(reader, &baseArray[i])
		}

		buildTree(1, 1, n)

		for query := 0; query < q; query++ {
			var op int
			fmt.Fscan(reader, &op)

			if op == 1 {
				var p int
				var v int64
				fmt.Fscan(reader, &p, &v)
				updateTree(1, 1, n, p, v)
			} else {
				var l, r int
				var m int64
				fmt.Fscan(reader, &l, &r, &m)

				resNode := queryTree(1, 1, n, l, r)

				trueTotalFixed := resNode.totalFixed
				U := resNode.totalFree

				if m < trueTotalFixed {
					fmt.Fprintln(writer, 0)
					continue
				}

				M := (m - trueTotalFixed + MOD) % MOD

				sum1 := resNode.sumF2
				if U == 0 {
					if M == 0 {
						fmt.Fprintln(writer, sum1)
					} else {
						fmt.Fprintln(writer, 0)
					}
					continue
				}

				sum2 := (2*resNode.sumFC + resNode.sumC) % MOD
				sum3 := (resNode.sumC2 + resNode.sumC) % MOD

				ans := (nCr(int(M+U-1), int(M)) * sum1) % MOD
				ans = (ans + nCr(int(M+U-1), int(M-1))*sum2) % MOD
				ans = (ans + nCr(int(M+U-1), int(M-2))*sum3) % MOD

				fmt.Fprintln(writer, ans)
			}
		}
	}
}
