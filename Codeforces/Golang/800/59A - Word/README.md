## 59A - Word

### Description
Given a word, convert it entirely to uppercase if it contains more uppercase letters than lowercase letters.
Otherwise, convert it entirely to lowercase.

### Mathematical Logic
Count the number of uppercase and lowercase letters in the string.

Let:
- `upper_count` = number of uppercase characters
- `lower_count` = number of lowercase characters

Then:
- if `upper_count > lower_count`, print the uppercase version of the word
- otherwise, print the lowercase version of the word

### Complexity
- Time: `O(n)` - each character is processed once.
- Space: `O(n)` - a transformed string is produced.

### Implementations
- [Go solution](main.go)
