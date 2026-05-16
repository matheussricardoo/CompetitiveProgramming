## 2228E1 - Amanojaku and Sequence (Easy Version)

### Link
https://codeforces.com/contest/2228/problem/E1

### Statement (easy version)
For a sequence c of non-negative integers define f(c) = sum_{i=1..|c|} (sum_{j=1..i} c_j)^2. For a given integer sequence b with b_i ≥ −1 and an integer m ≥ 0, g(b,m) is the sum of f(c) over all sequences c that are "valid" for (b,m) (validity defined in the problem). Given an array a of length n with a_i ≥ −1 and q = 1 queries of the form: given l, r, m compute g([a_l, ..., a_r], m) modulo 998244353.

### Input
First line t the number of test cases. For each test case: n and q (1 ≤ n ≤ 3·10^5, q = 1), then the array a_1..a_n (−1 ≤ a_i ≤ 10^6). Then q lines with queries (type 2 queries: l r m).

### Output
For each query of the second type output g([a_l,...,a_r], m) modulo 998244353.

### Notes
This is the easy version where q = 1 and op = 2; see problem link for full formal definition and constraints.

### Implementation
- [Go solution](main.go)