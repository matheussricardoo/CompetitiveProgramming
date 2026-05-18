## 977A - Wrong Subtraction

### Description
Given an integer `n` and a number of operations `k`, repeat the following operation exactly `k` times:
- if `n` ends with `0`, divide it by `10`
- otherwise, subtract `1`

### Mathematical Logic
Simulate the operations one by one for `k` iterations.

For each step:
- if `n % 10 == 0`, then `n = n / 10`
- else `n = n - 1`

After all operations, print the final value of `n`.

### Complexity
- Time: `O(k)` - one operation per iteration.
- Space: `O(1)`.

### Implementations
- [Rust solution](src/main.rs)
