## 133A - HQ9+

### Description
Implement the esoteric HQ9+ language semantics for given input character: output the appropriate string depending on the command character.

### Mathematical Logic
- `H` prints `Hello, World!`.
- `Q` prints the source code itself.
- `9` prints the lyrics of `99 Bottles of Beer` (or required output snippet).
- `+` increments an accumulator (ignored if not requested in problem).

### Complexity
- Time: `O(1)` per input (outputs are fixed or bounded).
- Space: `O(n)` if returning source code (n = program size), otherwise `O(1)`.

### Implementations
- [Rust solution](src/main.rs)