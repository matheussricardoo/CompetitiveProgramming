## 2228B - Remilia Plays Soku

### Link
https://codeforces.com/contest/2228/problem/B

### Statement
There are n positions arranged in a circle (positions 1..n). Initially Reimu is at x1 and Remilia at x2 (x1 ≠ x2). Each second Remilia may move to an adjacent position or stay, and she can move at most k times in total. After observing Remilia's move, Reimu moves one step or stays. If after both actions they are at the same position, Reimu catches Remilia. Assuming optimal play (Reimu minimizes the catching time, Remilia maximizes it), compute the number of seconds until capture.

### Input
The first line contains t, the number of test cases. Each test case is a single line with four integers: n, x1, x2, k (2 ≤ n ≤ 1e8, 1 ≤ x1,x2 ≤ n, x1 ≠ x2, 0 ≤ k ≤ 1e8).

### Output
For each test case print the number of seconds until Reimu catches Remilia.

### Example
Input:
4
2 1 2 0
4 3 2 1
4 2 3 1
16 8 4 2

Output:
1
2
2
6

### Implementation
- [Go solution](main.go)