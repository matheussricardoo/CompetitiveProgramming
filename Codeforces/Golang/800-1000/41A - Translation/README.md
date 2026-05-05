## 41A - Translation

### Description
Given two strings `s` and `t`, determine whether `t` is the reverse of `s`.

Print `YES` if `t` equals the reverse of `s`, otherwise print `NO`.

### Mathematical Logic
Compute `reverse(s)` and compare with `t`:
- if `reverse(s) == t`, print `YES`
- otherwise, print `NO`

### Complexity
- Time: `O(n)` - reversing and comparing both strings of length `n`.
- Space: `O(n)` - for the reversed string.

### Implementations
- [Go solution](main.go)
