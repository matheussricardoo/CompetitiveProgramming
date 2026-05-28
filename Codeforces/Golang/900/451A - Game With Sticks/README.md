## 451A - Game With Sticks

### Description
Two players play on an n×m board removing a 1×1 square each turn; the player unable to move loses. Players alternate turns removing exactly one square from the board. Determine who wins assuming optimal play.

### Mathematical Logic
Each move removes one square; the total number of moves possible is n*m. If n*m is odd, the first player wins; otherwise the second player wins. Equivalently, the winner is determined by parity of min(n,m): if min(n,m) is odd, "Akshat" wins; else "Malvika" wins (based on the board removal game symmetry and problem statement interpretation).

### Complexity
- Time: `O(1)`.
- Space: `O(1)`.

### Implementations
- [Go solution](main.go)