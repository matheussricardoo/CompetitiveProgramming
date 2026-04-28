## 112A - Petya and Strings

### Description
The goal is to compare two strings lexicographically while ignoring letter case.

- Print `-1` if the first string is smaller.
- Print `1` if the first string is greater.
- Print `0` if both strings are equal.

### Mathematical Logic
Case-insensitive comparison can be reduced to a normal lexicographic comparison by converting both strings to the same case.

Let:
- `a = lowercase(s1)`
- `b = lowercase(s2)`

Then:
- if `a < b`, answer is `-1`
- if `a > b`, answer is `1`
- otherwise, answer is `0`

### Complexity
- Time: `O(n)` - each character is processed at most once for normalization/comparison.
- Space: `O(n)` - lowercased copies of the strings are created.

### Implementations
- [Rust solution](src/main.rs)
