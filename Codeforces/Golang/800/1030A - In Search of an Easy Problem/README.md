## 1030A - In Search of an Easy Problem

### Description
Given a list of answers to a problem set, determine whether at least one participant said `1`.

- If there is at least one `1`, print `HARD`
- Otherwise, print `EASY`

### Mathematical Logic
Check whether the input contains the value `1`.

If any number equals `1`, the problem is hard.
Otherwise, it is easy.

### Complexity
- Time: `O(n)` - each answer is inspected once.
- Space: `O(1)`.

### Implementations
- [Go solution](main.go)
