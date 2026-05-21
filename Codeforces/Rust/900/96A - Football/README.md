## 96A - Football

### Description
Given a string of characters `0` and `1`, determine whether it contains a substring of at least five equal consecutive characters (`00000` or `11111`).

### Mathematical Logic
- Scan the string and track run length of consecutive equal characters.
- If any run length ≥ 5, print `YES`, otherwise `NO`.

### Complexity
- Time: `O(n)`.
- Space: `O(1)`.

### Implementations
- [Rust solution](src/main.rs)