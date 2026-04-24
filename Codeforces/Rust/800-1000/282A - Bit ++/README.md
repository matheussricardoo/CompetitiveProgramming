## 282A - Bit++

### Description
The goal is to compute the final value of variable `X` after executing `n` statements.

Each statement is one of:
- `++X`
- `X++`
- `--X`
- `X--`

The variable starts at `0`.

### Mathematical Logic
Initialize `X = 0`.
For each instruction:
- If it represents increment (`++X` or `X++`), add `1` to `X`.
- Otherwise, subtract `1` from `X`.

After processing all instructions, print `X`.

### Complexity
- Time: `O(n)` - each statement is processed once.
- Space: `O(1)` - only a few scalar variables are used.

### Implementations
- [Rust solution](src/main.rs)
