## 546A - Soldier and Bananas

### Description
A soldier wants to buy `w` bananas. The price of the first banana is `k`, the second is `2k`, the third is `3k`, and so on.

The goal is to compute how much money he needs to borrow from his friend if he already has `n` dollars.

### Mathematical Logic
The total cost of `w` bananas is the sum of an arithmetic progression:

- `k + 2k + 3k + ... + wk`

This can be written as:

- `k * (w * (w + 1) / 2)`

Then:
- `borrowed = total_cost - n`
- if `borrowed < 0`, print `0`
- otherwise, print `borrowed`

### Complexity
- Time: `O(1)` - only arithmetic operations are used.
- Space: `O(1)`.

### Implementations
- [Go solution](main.go)
