## 405A - Gravity Flip

### Description
Given a row of blocks with colors denoted by integers, simulate gravity flip which sorts the blocks by color increasing when the box is flipped.

### Mathematical Logic
- After flip, the blocks settle so their colors are in non-decreasing order; simply sort the array.

### Complexity
- Time: `O(n log n)` for sorting.
- Space: `O(1)` additional.

### Implementations
- [Go solution](main.go)
