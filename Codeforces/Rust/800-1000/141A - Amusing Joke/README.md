## 141A - Amusing Joke

### Description
Given two strings `s1` and `s2`, and a third string `pile`, determine if `pile` is formed by combining all characters from `s1` and `s2`.

### Mathematical Logic
Concatenate `s1` and `s2`, then compare the sorted characters with the sorted characters of `pile`.

If both multisets of characters are identical, print `YES`; otherwise, print `NO`.

### Complexity
- Time: `O(n log n)` - due to sorting the characters.
- Space: `O(n)`.

### Implementations
- [Rust solution](src/main.rs)
