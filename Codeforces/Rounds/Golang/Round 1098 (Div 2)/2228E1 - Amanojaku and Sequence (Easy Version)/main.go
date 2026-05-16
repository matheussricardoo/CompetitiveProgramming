// Problem: E1. Amanojaku and Sequence (Easy Version)
// Link: https://codeforces.com/contest/2228/problem/E1


package main

import (
	"bufio"
	"fmt"
	"os"
)

const MOD = 998244353

var fact []int64
var invFact []int64

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

func getSum(pref []int64, l, r int) int64 {
	val := (pref[r] - pref[l-1]) % MOD
	if val < 0 {
		val += MOD
	}
	return val
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	precompute(1400000)

	var t int
	fmt.Fscan(reader, &t)

	for tc := 0; tc < t; tc++ {
		var n, q int
		fmt.Fscan(reader, &n, &q)

		a := make([]int64, n+1)
		AF := make([]int64, n+1)
		AC := make([]int64, n+1)

		prefAF := make([]int64, n+1)
		prefAF2 := make([]int64, n+1)
		prefAC := make([]int64, n+1)
		prefAC2 := make([]int64, n+1)
		prefAFAC := make([]int64, n+1)

		for i := 1; i <= n; i++ {
			fmt.Fscan(reader, &a[i])

			fixedVal := int64(0)
			isFree := int64(0)
			if a[i] >= 0 {
				fixedVal = a[i]
			} else {
				isFree = 1
			}

			AF[i] = (AF[i-1] + fixedVal) % MOD
			AC[i] = (AC[i-1] + isFree) % MOD

			prefAF[i] = (prefAF[i-1] + AF[i]) % MOD
			prefAF2[i] = (prefAF2[i-1] + (AF[i]*AF[i])%MOD) % MOD
			prefAC[i] = (prefAC[i-1] + AC[i]) % MOD
			prefAC2[i] = (prefAC2[i-1] + (AC[i]*AC[i])%MOD) % MOD
			prefAFAC[i] = (prefAFAC[i-1] + (AF[i]*AC[i])%MOD) % MOD
		}
		for query := 0; query < q; query++ {
			var op, l, r int
			var m int64
			fmt.Fscan(reader, &op, &l, &r, &m)

			totalFixed := (AF[r] - AF[l-1] + MOD) % MOD
			U := (AC[r] - AC[l-1] + MOD) % MOD
			M := (m - totalFixed + MOD) % MOD
			if m < totalFixed {
				fmt.Fprintln(writer, 0)
				continue
			}
			t1_1 := getSum(prefAF2, l, r)
			t1_2 := (2 * AF[l-1] % MOD) * getSum(prefAF, l, r) % MOD
			t1_3 := (AF[l-1] * AF[l-1] % MOD) * int64(r-l+1) % MOD
			sum1 := (t1_1 - t1_2 + t1_3 + 2*MOD) % MOD

			if U == 0 {
				if M == 0 {
					fmt.Fprintln(writer, sum1)
				} else {
					fmt.Fprintln(writer, 0)
				}
				continue
			}

			t2_1 := 2 * getSum(prefAFAC, l, r) % MOD
			t2_2 := getSum(prefAC, l, r)
			t2_3 := (2 * AC[l-1] % MOD) * getSum(prefAF, l, r) % MOD
			t2_4 := (2 * AF[l-1] % MOD) * getSum(prefAC, l, r) % MOD
			t2_5 := (2 * AF[l-1] % MOD * AC[l-1] % MOD) * int64(r-l+1) % MOD
			t2_6 := AC[l-1] * int64(r-l+1) % MOD
			sum2 := (t2_1 + t2_2 - t2_3 - t2_4 + t2_5 - t2_6 + 4*MOD) % MOD

			t3_1 := getSum(prefAC2, l, r)
			t3_2 := getSum(prefAC, l, r)
			t3_3 := (2 * AC[l-1] % MOD) * getSum(prefAC, l, r) % MOD
			t3_4 := (AC[l-1] * AC[l-1] % MOD) * int64(r-l+1) % MOD
			t3_5 := AC[l-1] * int64(r-l+1) % MOD
			sum3 := (t3_1 + t3_2 - t3_3 + t3_4 - t3_5 + 3*MOD) % MOD

			ans := (nCr(int(M+U-1), int(M)) * sum1) % MOD
			ans = (ans + nCr(int(M+U-1), int(M-1))*sum2) % MOD
			ans = (ans + nCr(int(M+U-1), int(M-2))*sum3) % MOD

			fmt.Fprintln(writer, ans)
		}
	}
}
