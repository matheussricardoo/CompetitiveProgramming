## 160A - Twins

### Description
Given an array of coin values, choose the minimum number of largest-valued coins such that their sum strictly exceeds the sum of the remaining coins.

### Mathematical Logic
- Sort coin values descending.
- Take coins from largest to smallest until chosen_sum > total_sum - chosen_sum.
- Return the count of chosen coins.

### Complexity
- Time: `O(n log n)` due to sorting.
- Space: `O(1)` additional.

### Implementations
- [Rust solution](src/main.rs)