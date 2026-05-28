## 1374B - Multiply by 2, divide by 6

### Description
Given an integer n, you are allowed to perform operations: if n is divisible by 6, replace n with n/6; else you may multiply n by 2. Determine the minimum number of operations required to transform n into 1, or output -1 if impossible.

### Mathematical Logic
Factor n into powers of 2 and 3: n = 2^{c2} * 3^{c3} * remaining. If the remaining part ≠ 1, it is impossible. Each division by 6 reduces both c3 and c2 by 1; multiplication by 2 increases c2 by 1. To reach 1, we need c2 ≤ c3, otherwise impossible. The minimal operations equal (c3 - c2) (number of times we must multiply by 2 to balance) + c3 (number of divide-by-6 operations).

### Complexity
- Time: `O(log n)` per test case (counting factors).
- Space: `O(1)`.

### Implementations
- [Go solution](main.go)