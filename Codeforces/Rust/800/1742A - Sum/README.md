## 1742A - Sum

### Description
For each test case, determine if two of the three integers sum to the third.

Print `YES` if any pair sums to the remaining number, otherwise print `NO`.

### Mathematical Logic
For numbers `a`, `b`, and `c`, check:
- `a + b == c`
- `a + c == b`
- `b + c == a`

If any condition is true, the answer is `YES`.

### Complexity
- Time: `O(1)` per test case.
- Space: `O(1)`.

### Implementations
- [Rust solution](src/main.rs)
