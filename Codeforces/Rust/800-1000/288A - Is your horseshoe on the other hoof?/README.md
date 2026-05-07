## 288A - Is your horseshoe on the other hoof?

### Description
Given four integers representing the colors of four horseshoes, determine how many horseshoes must be replaced so that all four are of different colors.

### Mathematical Logic
Place the four colors in a set to count distinct colors. The number to replace is `4 - distinct_count`.

### Complexity
- Time: `O(1)` - fixed number of inputs.
- Space: `O(1)`.

### Implementations
- [Rust solution](src/main.rs)
