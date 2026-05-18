## 510A - Fox And Snake

### Description
Draw an `n x m` grid with a snake pattern:
- Odd rows: all `#`
- Even rows: `#` at one end, `.` elsewhere, alternating which end

### Mathematical Logic
For row `i` (1-indexed):
- if odd: print all `#`
- if even:
  - if row 2, 6, 10, ... (i % 4 == 2): `#` at the right
  - if row 4, 8, 12, ... (i % 4 == 0): `#` at the left

### Complexity
- Time: `O(n * m)` - generating the grid.
- Space: `O(n * m)` - storing the output.

### Implementations
- [Rust solution](src/main.rs)
