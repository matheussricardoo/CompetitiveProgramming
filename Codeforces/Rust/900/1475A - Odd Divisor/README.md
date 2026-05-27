## 1475A - Odd Divisor

### Description
Given an integer n, determine whether n has an odd divisor greater than 1 (equivalently, whether n is not a power of two). For each test case print `YES` if such an odd divisor exists, otherwise print `NO`.

### Mathematical Logic
An integer n has no odd divisors greater than 1 if and only if n is a power of two. Use the bit trick `n & (n - 1) == 0` to test for powers of two. Output `YES` when `n & (n - 1) != 0`.

### Complexity
- Time: `O(1)` per test case (bitwise operations).
- Space: `O(1)`.

### Implementations
- [Rust solution](src/main.rs)