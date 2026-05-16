## 2228E2 - Amanojaku and Sequence (Hard Version)

### Link
https://codeforces.com/contest/2228/problem/E2

### Statement (hard version)
As in E1, but the harder version supports q up to 3·10^5 and two types of operations: point updates (set a_p = v) and queries that compute g([a_l,...,a_r], m) modulo 998244353. You need a data structure supporting updates and range queries efficiently.

### Input
The first line contains t, the number of test cases. Each test case contains n and q (1 ≤ n ≤ 3·10^5, 1 ≤ q ≤ 3·10^5), the array a_1..a_n (−1 ≤ a_i ≤ 10^6), then q queries. Each query is either:
- `1 p v`: set a_p = v (−1 ≤ v ≤ 10^6), or
- `2 l r m`: compute g([a_l,...,a_r], m) modulo 998244353.

### Output
For each type-2 query output the answer modulo 998244353.

### Implementation notes
Use precomputed factorials and a segment tree where each node stores aggregated statistics needed to compute contributions for a segment.

### Implementation
- [Go solution](main.go)