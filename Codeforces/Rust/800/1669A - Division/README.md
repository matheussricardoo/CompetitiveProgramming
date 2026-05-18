## 1669A - Division

### Description
Given a Codeforces user's rating, determine their division level:
- Division 1: rating ≥ 1900
- Division 2: 1600 ≤ rating ≤ 1899
- Division 3: 1400 ≤ rating ≤ 1599
- Division 4: rating < 1400

### Mathematical Logic
Simply compare the input rating against the division thresholds in the following order:
1. Check if rating ≥ 1900 → Division 1
2. Else check if rating ≥ 1600 → Division 2
3. Else check if rating ≥ 1400 → Division 3
4. Else → Division 4

### Complexity
- Time: `O(1)` - constant time comparison.
- Space: `O(1)`.

### Implementations
- [Rust solution](src/main.rs)
