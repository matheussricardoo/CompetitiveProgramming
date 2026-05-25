## 337A - Puzzles

### Description
Given `n` students and `m` puzzle pieces with different sizes, choose exactly `n` pieces such that the difference between the largest and smallest chosen piece is minimized.

### Mathematical Logic
- Sort the puzzle sizes in non-decreasing order.
- Consider every contiguous subarray of length `n`.
- For each subarray, compute `max - min` and keep the minimum value.

### Complexity
- Time: `O(m log m)` - sorting dominates.
- Space: `O(1)` additional.

### Implementations
- [Rust solution](src/main.rs)
