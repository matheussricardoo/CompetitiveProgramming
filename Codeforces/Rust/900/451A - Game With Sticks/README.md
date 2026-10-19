## 451A - Game With Sticks

### Description
Two players alternately choose an intersection and remove its horizontal and vertical sticks; the player who cannot move loses. Determine the winner assuming both play optimally.

### Mathematical Logic
Each turn removes one horizontal and one vertical stick, so the number of moves is min(n,m): if min(n,m) is odd, the first player (Akshat) wins; otherwise the second player (Malvika) wins.

### Complexity
- Time: `O(1)`.
- Space: `O(1)`.

### Implementations
- [Rust solution](src/main.rs)