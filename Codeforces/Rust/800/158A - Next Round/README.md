## 158A - Next Round

### Description
The goal is to count how many participants advance to the next round.

A participant advances if:
- their score is at least the score of the participant in position `k`, and
- their score is greater than `0`.

### Mathematical Logic
Given scores in non-increasing order:
- Let `threshold = a[k - 1]`.
- Count all scores `a[i]` such that:
  - `a[i] >= threshold`
  - `a[i] > 0`

The final count is the number of qualified participants.

### Complexity
- Time: `O(n)` - each score is checked at most once.
- Space: `O(1)` extra space (excluding input storage).

### Implementations
- [Rust solution](src/main.rs)
