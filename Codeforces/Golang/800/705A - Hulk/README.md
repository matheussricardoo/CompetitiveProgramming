## 705A - Hulk

### Description
Hulk expresses his feelings in a pattern: he says `I hate`, then `I love`, alternating for `n` clauses, joined by `that`, and ending with `it`.

Example for `n=3`: `I hate that I love that I hate it`.

### Mathematical Logic
Build a sequence of `n` phrases where odd positions are `I hate` and even positions are `I love`, joined by ` that ` and finished with ` it`.

### Complexity
- Time: `O(n)` - concatenating `n` parts.
- Space: `O(n)`.

### Implementations
- [Go solution](main.go)
