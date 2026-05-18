## 339A - Helpful Maths

### Description
The input is a sum of single-digit numbers (`1`, `2`, `3`) separated by `+` signs, like `3+2+1`.

The goal is to rearrange the numbers in non-decreasing order and print them again in the same `+`-separated format.

### Mathematical Logic
Let the expression be split into terms:

- `terms = split(s, '+')`

Then:
- sort `terms` in ascending order
- join them back with `+`

This produces the lexicographically and numerically smallest valid expression with the same digits.

### Complexity
- Time: `O(n log n)` - sorting dominates, where `n` is the number of terms.
- Space: `O(n)` - to store the split terms.

### Implementations
- [Go solution](main.go)
