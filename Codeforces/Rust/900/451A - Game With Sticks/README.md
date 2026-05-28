## 451A - Game With Sticks

### Description
Two players alternately remove a 1×1 square from an n×m grid; the player who cannot move loses. Determine the winner assuming both play optimally.

### Mathematical Logic
The winner is determined by the parity of the number of moves (n*m), or equivalently by parity of min(n,m) for this problem's specific interpretation: if min(n,m) is odd, the first player (Akshat) wins; otherwise the second player (Malvika) wins.

### Complexity
- Time: `O(1)`.
- Space: `O(1)`.

### Implementations
- [Rust solution](src/main.rs)