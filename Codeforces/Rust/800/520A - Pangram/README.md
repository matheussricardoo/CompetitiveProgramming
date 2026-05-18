## 520A - Pangram

### Description
A pangram is a string that contains all 26 letters of the English alphabet.

Given a string, determine whether it is a pangram (print `YES` or `NO`).

### Mathematical Logic
Convert the string to lowercase. Collect all distinct lowercase letters in a set. If the set size equals `26`, it is a pangram.

### Complexity
- Time: `O(n)` - iterate through all characters once.
- Space: `O(1)` - at most 26 letters in the set.

### Implementations
- [Rust solution](src/main.rs)
