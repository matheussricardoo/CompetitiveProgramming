## 4A - Watermelon

### Description
The goal is to determine whether a watermelon of weight `w` can be split into two positive even parts.

### Mathematical Logic
To split `w` into two positive even integers, two conditions must hold:
- `w` must be even, because the sum of two even numbers is always even.
- `w` must be greater than `2`, since the smallest valid split is `2 + 2 = 4`.

### Complexity
- Time: `O(1)` - only one arithmetic check is needed.
- Space: `O(1)` - no extra data structures are used.

### Implementations
- [Rust solution](src/main.rs)
