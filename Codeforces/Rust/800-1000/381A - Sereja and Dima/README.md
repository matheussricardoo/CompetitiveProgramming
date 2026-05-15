## 381A - Sereja and Dima

### Description
Sereja and Dima play a card game with `n` cards laid in a row, each containing a distinct integer.
They alternate turns (Sereja goes first), and on each turn a player must take either the leftmost or the rightmost card.
Both players are greedy: each always picks the card with the larger value.

Given the initial row of cards, determine each player's final score.

### Mathematical Logic
Use two pointers `l = 0` and `r = n - 1` to represent the current ends of the row.

On each turn `t` (0-indexed):
- The current player picks `max(a[l], a[r])`
- If `a[l] > a[r]`: take `a[l]`, advance `l++`
- Else: take `a[r]`, retreat `r--`
- If `t` is even → add to Sereja's score; otherwise → add to Dima's score

Repeat until `l > r`.

Since all values are distinct, there is never a tie between the two ends, so the greedy choice is always well-defined.

### Complexity
- Time: `O(n)` - a single left-to-right sweep with two pointers.
- Space: `O(n)` - to store the input array.

### Implementations
- [Rust solution](src/main.rs)