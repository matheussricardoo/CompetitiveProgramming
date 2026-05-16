## 2228F - Momoyo and the Network

### Link
https://codeforces.com/contest/2228/problem/F

### Statement
Given a tree with n nodes where node i has weight a_i, select a simple path of exactly k edges and remove all edges on it. That splits the tree into k+1 connected components; each component has a weight equal to the sum of its nodes' weights. Maximize the minimum component weight after removing such a path, or output −1 if no simple path of exactly k edges exists.

### Input
The first line contains t, the number of test cases. Each test case begins with n and k (1 ≤ k ≤ n−1, 2 ≤ n ≤ 2·10^5), followed by a line with n integers a_1..a_n (1 ≤ a_i ≤ 10^9), then n−1 lines with edges u v describing the tree. It is guaranteed the sum of n over test cases does not exceed 2·10^5.

### Output
For each test case print the maximum possible minimum component weight, or −1 if no simple path of length k exists.

### Example
See the problem link for sample tests.

### Implementation notes
- Compute subtree sums and a root-to-leaves order. For a candidate X use a DP that computes contributions from children and validates whether enough components with sum ≥ X can be formed; use binary search on X over [1..total_sum].

### Implementation
- [Go solution](main.go)