## 208A - Dubstep

### Description
Given a song encoded with the string `WUB`, restore the original words by removing the noise and separating words with a single space.

### Mathematical Logic
Replace every occurrence of `WUB` with a space, then collapse multiple spaces into one and trim leading or trailing spaces.

### Complexity
- Time: `O(n)` - scanning and replacing the input string of length `n`.
- Space: `O(n)` - for the transformed string and split words.

### Implementations
- [Rust solution](src/main.rs)
