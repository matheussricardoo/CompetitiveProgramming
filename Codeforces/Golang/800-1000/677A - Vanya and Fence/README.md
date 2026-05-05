## 677A - Vanya and Fence

### Description
Vanya wants to walk along a fence of `n` fence sections. Each section has a height. Vanya can walk along a section if its height is less than or equal to his height `h`; otherwise he must bend and that section contributes `2` to the required width. Sections with height `<= h` contribute `1`.

Calculate the total width Vanya needs to pass the fence, summing `1` for short sections and `2` for tall sections.

### Mathematical Logic
For each fence section with height `a[i]`:
- add `1` to the total if `a[i] <= h`
- add `2` if `a[i] > h`

Total width = sum over sections of (1 if `a[i] <= h` else 2).

### Complexity
- Time: `O(n)` - single pass over the heights.
- Space: `O(1)`.

### Implementations
- [Go solution](main.go)
