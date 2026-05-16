## 2228C1 - Cirno and Number (Easy Version)

### Link
https://codeforces.com/contest/2228/problem/C1

### Statement
(Easy version, n = 2) Given a non-negative integer a and an increasing sequence of digits d of length n (n = 2), find the minimum value of |a − b| over all non-negative integers b whose decimal representation contains only digits from d.

### Input
The first line contains t, the number of test cases. Each test case contains a line with `a` (0 ≤ a ≤ 10^17) and `n` (=2), followed by a line with the digits d_1, d_2 (0 ≤ d_1 < d_2 ≤ 9).

### Output
For each test case print the minimum value of |a − b|.

### Example
Input:
4
0 2
0 2
11 2
1 2
222 2
3 4
3333 2
6 7

Output:
0
0
111
2556

### Implementation
- [Go solution](main.go)