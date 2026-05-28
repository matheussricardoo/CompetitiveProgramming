## 1374B - Multiply by 2, divide by 6

### Description
Given an integer n, you may perform operations: if n is divisible by 6, replace n with n/6; otherwise you may multiply n by 2. Determine the minimum number of operations to reach 1, or output -1 if impossible.

### Mathematical Logic
Write n = 2^{c2} * 3^{c3} * rem. If rem ≠ 1, it's impossible. We must have c2 ≤ c3; otherwise we cannot eliminate extra factors of 2 by dividing by 6. The minimal operations are (c3 - c2) multiplications by 2 to increase c2, plus c3 divisions by 6: total (c3 - c2) + c3.

### Complexity
- Time: `O(log n)` per test case (factor counting).
- Space: `O(1)`.

### Implementations
- [Rust solution](src/main.rs)