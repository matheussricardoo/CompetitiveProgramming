## 1328A - Divisibility Problem

### Description
Given integers `a` and `b`, find the minimum number of times we must increase `a` by 1 to make it divisible by `b`.

### Mathematical Logic
- If `a % b == 0`, no increment is needed, result is `0`.
- Otherwise, the remainder is `a % b`, and we need `b - (a % b)` increments to reach the next multiple of `b`.

### Complexity
- Time: `O(1)` - constant arithmetic.
- Space: `O(1)`.

### Implementations
- [Rust solution](src/main.rs)
