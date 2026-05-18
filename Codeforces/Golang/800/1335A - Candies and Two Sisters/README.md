## 1335A - Candies and Two Sisters

### Description
Two sisters want to divide `n` candies fairly. They divide by repeatedly giving one candy at a time to each sister. Find how many candies the second sister gets when the first sister stops (after she gets the first candy).

### Mathematical Logic
In a fair division with alternation:
- First sister gets candies `1, 3, 5, ..., 2k-1` (odd positions)
- Second sister gets candies `2, 4, 6, ..., 2j` (even positions)

If `n` is even, both get `n/2`. If `n` is odd, the first gets one more.

Result for second sister = `(n - 1) / 2` (integer division).

### Complexity
- Time: `O(1)` - constant arithmetic.
- Space: `O(1)`.

### Implementations
- [Go solution](main.go)
