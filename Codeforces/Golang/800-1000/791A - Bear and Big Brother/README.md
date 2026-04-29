## 791A - Bear and Big Brother

### Description
Given two integers `a` and `b` representing two bears' weights, determine how many years it will take for `a` to become strictly greater than `b` if every year `a` triples and `b` doubles.

### Mathematical Logic
Simulate years until `a > b`:
- each year: `a = a * 3`, `b = b * 2`.
- count the number of iterations required.

### Complexity
- Time: `O(k)` - `k` is the number of years until `a > b` (small for problem constraints).
- Space: `O(1)`.

### Implementations
- [Go solution](main.go)
