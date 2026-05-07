## 136A - Presents

### Description
Given a permutation `p` of numbers from `1` to `n` where `p[i]` denotes the friend who received the `i`-th present, compute the inverse permutation where position `j` contains the index of the present given to friend `j`.

### Mathematical Logic
For each position `i` (1-based), let friend `p[i]` receive present `i`. Build an array `result` such that `result[p[i]] = i`.

### Complexity
- Time: `O(n)` - single pass to build the inverse mapping.
- Space: `O(n)` - to store the result.

### Implementations
- [Rust solution](src/main.rs)
