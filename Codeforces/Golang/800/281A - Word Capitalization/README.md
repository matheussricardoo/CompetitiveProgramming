## 281A - Word Capitalization

### Description
The goal is to capitalize the first letter of a single-word string.

Given a lowercase word, transform its first character to uppercase and print the resulting word.

### Mathematical Logic
- Let `s` be the input word.
- If `s` is empty, output an empty string.
- Otherwise, output `uppercase(s[0]) + s[1:]`.

### Complexity
- Time: `O(n)` - converting and copying characters once.
- Space: `O(n)` - a new string is produced.

### Implementations
- [Go solution](main.go)
