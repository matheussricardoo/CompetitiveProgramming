## 263A - Beautiful Matrix

### Description
The goal is to find the minimum number of moves needed to bring the `1` in a `5 x 5` matrix to the center cell.

A single move consists of swapping two adjacent rows or columns, which is equivalent to moving the `1` one step up, down, left, or right.

### Mathematical Logic
The center of the matrix is at position `(3, 3)` using 1-based indexing.

If the `1` is currently at position `(i, j)`, the minimum number of moves is the Manhattan distance to the center:

- `|i - 3| + |j - 3|`

This works because each move changes only one coordinate by `1`.

### Complexity
- Time: `O(25)` - the matrix has a fixed size of `5 x 5`.
- Space: `O(1)` - only a few variables are used.

### Implementations
- [Go solution](main.go)
