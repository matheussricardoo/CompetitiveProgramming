## 61A - Ultra-Fast Mathematician

### Description
Given two binary strings of equal length, produce a new string where each bit is `1` if the corresponding bits differ and `0` if they are the same (bitwise XOR).

### Mathematical Logic
For each position `i`:
- result[i] = `1` if `a[i] != b[i]`, else `0`.

This is equivalent to bitwise XOR for binary digits.

### Complexity
- Time: `O(n)` — process each character once.
- Space: `O(n)` — to build the result string.

### Implementations
- [Rust solution](src/main.rs)
