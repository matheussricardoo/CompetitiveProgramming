## 71A - Way Too Long Words

### Description
The goal is to process each word and abbreviate it when its length is greater than `10`.

### Mathematical Logic
For each word `s`:
- If `len(s) <= 10`, print the word unchanged.
- Otherwise, print:
  - the first character,
  - the number of removed middle characters (`len(s) - 2`),
  - the last character.

This produces abbreviations like `localization -> l10n` and `internationalization -> i18n`.

### Complexity
- Time: `O(n * m)` - where `n` is the number of words and `m` is the average word length.
- Space: `O(1)` extra space per word (excluding input storage).

### Implementations
- [Go solution](main.go)
