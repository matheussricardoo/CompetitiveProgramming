## 2228D - Sanae, Cross and Color

### Link
https://codeforces.com/contest/2228/problem/D

### Statement
Given n distinct integer points (x_i, y_i) with 1 ≤ x_i, y_i ≤ n, choose integers k1 and k2 such that each of the four regions divided by the lines x = k1+0.5 and y = k2+0.5 contains at least one point. Color points according to the region they lie in. Two colorings are distinct if there exists at least one point colored differently. Count the number of distinct colorings possible.

### Input
The first line contains t, the number of test cases. Each test case starts with n (4 ≤ n ≤ 2e6), followed by n lines with two integers x_i y_i (1 ≤ x_i, y_i ≤ n). It is guaranteed points are pairwise distinct and sum of n over tests ≤ 2e6.

### Output
For each test case output the number of distinct colorings.

### Example
No example included here; see problem link for samples.

### Implementation
- [Go solution](main.go)