## 617A - Elephant

### Description
Given the number of steps `n` the elephant needs to move on the x-axis, determine the minimum number of moves where in one move the elephant can walk `1`, `2`, `3` or `4` or `5` steps. The optimal move is to use as many `5`-step moves as possible.

### Mathematical Logic
Minimum moves = `ceil(n / 5)`.

### Complexity
- Time: `O(1)` - constant arithmetic.
- Space: `O(1)`.

### Implementations
- [Rust solution](src/main.rs)
