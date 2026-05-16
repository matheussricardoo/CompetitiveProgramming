## 2228A - Marisa Steals Reimu's Takeout

### Link
https://codeforces.com/contest/2228/problem/A

### Statement
Marisa has a sequence w of length n where each element is 0, 1 or 2. She can repeat the following operation: select a non-empty subsequence of w whose sum is divisible by 3 and remove its elements. Determine the maximum number of operations she can perform.

### Input
The first line contains t (1 ≤ t ≤ 500), the number of test cases. Each test case consists of:
- an integer n (1 ≤ n ≤ 100)
- a line with n integers w_1, w_2, ..., w_n (0 ≤ w_i ≤ 2)

### Output
For each test case print the maximum number of operations Marisa can perform.

### Example
Input:
3
4
0 0 0 0
3
1 2 0
5
1 2 1 2 1

Output:
4
2
2

### Implementation
- [Go solution](main.go)