## 469A - I Wanna Be the Guy

### Description
The guy wants to beat a video game with `n` levels. Two friends provide lists of levels they can beat.

The goal is to determine if together they can beat all `n` levels (collect all from both lists and check if it covers `1` to `n`).

### Mathematical Logic
Collect all distinct levels from both friends' lists into a set. If the set size equals `n`, they can beat all levels.

### Complexity
- Time: `O(n)` - inserting all levels into the set.
- Space: `O(n)` - storing the set.

### Implementations
- [Rust solution](src/main.rs)
