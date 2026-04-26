## 50A - Domino piling

### Description
The goal is to place the maximum number of `2 x 1` dominoes on an `m x n` board.

Each domino covers exactly `2` cells, and dominoes cannot overlap.

### Mathematical Logic
- The board has `m * n` total cells.
- Since each domino covers `2` cells, the maximum number of dominoes is:
  - `(m * n) / 2` using integer division.

This automatically handles odd numbers of cells by discarding the leftover cell.

### Complexity
- Time: `O(1)` - only one arithmetic computation is performed.
- Space: `O(1)` - only scalar variables are used.

### Implementations
- [Rust solution](src/main.rs)
