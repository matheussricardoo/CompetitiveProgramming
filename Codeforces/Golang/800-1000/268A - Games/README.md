## 268A - Games

### Description
Given the home and away uniforms of `n` teams, count how many matches will have both teams wearing the same color.

A match is counted when the home color of one team matches the away color of another team.

### Mathematical Logic
For every pair of teams `(i, j)` with `i != j`, count a match if `home[i] == away[j]`.

### Complexity
- Time: `O(n^2)` - all pairs are checked.
- Space: `O(n)`.

### Implementations
- [Go solution](main.go)
