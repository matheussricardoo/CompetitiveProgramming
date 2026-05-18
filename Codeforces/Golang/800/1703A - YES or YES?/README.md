## 1703A - YES or YES?

### Description
For each test case, determine whether the given string is equal to `yes` in any combination of uppercase/lowercase letters.

Print `YES` if the string matches case-insensitively, otherwise print `NO`.

### Mathematical Logic
Convert the input string to lowercase and compare it with `yes`.

- if equal, print `YES`
- otherwise, print `NO`

### Complexity
- Time: `O(n)` - where `n` is the string length.
- Space: `O(n)`.

### Implementations
- [Go solution](main.go)
