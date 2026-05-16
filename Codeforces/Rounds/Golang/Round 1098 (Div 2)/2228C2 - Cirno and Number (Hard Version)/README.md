## 2228C2 - Cirno and Number (Hard Version)

### Link
https://codeforces.com/contest/2228/problem/C2

### Statement
Hard version: given a non-negative integer a and a strictly increasing sequence of digits d of length n (1 ≤ n ≤ 10), find the minimum value of |a − b| over all non-negative integers b whose decimal representation contains only digits from d.

### Input
The first line contains t, the number of test cases. Each test case contains `a` and `n` (0 ≤ a ≤ 10^17, 1 ≤ n ≤ 10), followed by a line with n distinct digits d_1 < ... < d_n (0 ≤ d_i ≤ 9).

### Output
For each test case print the minimum value of |a − b|.

### Example
Input:
4
0 1
0
11 2
1 2
222 3
2 3 4
3333 4
6 7 8 9

Output:
0
0
111
2334

### Implementation
- [Go solution](main.go)