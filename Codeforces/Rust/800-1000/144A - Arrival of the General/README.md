## 144A - Arrival of the General

### Description
Given an array of heights representing soldiers, find the minimum number of swaps needed to place the tallest soldier first and the shortest soldier last.

### Mathematical Logic
Find indices:
- `max_idx` = position of tallest soldier (first occurrence)
- `min_idx` = position of shortest soldier (last occurrence)

Moves to bring max to front: `max_idx`
Moves to bring min to back: `n - 1 - min_idx`

Total = `max_idx + (n - 1 - min_idx)`

If `max_idx > min_idx`, subtract 1 (they shift during same operation).

### Complexity
- Time: `O(n)` - single pass to find indices.
- Space: `O(1)`.

### Implementations
- [Rust solution](src/main.rs)
