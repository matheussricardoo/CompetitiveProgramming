## 231A - Team

### Description
The goal is to count how many problems the team will solve.

For each problem, three teammates indicate confidence (`0` or `1`).
The team solves a problem if at least two teammates are sure about the solution.

### Mathematical Logic
For each problem, read values `p`, `v`, and `t`:
- Compute `p + v + t`.
- If the sum is at least `2`, increment the answer.

At the end, print the total count of solvable problems.

### Complexity
- Time: `O(n)` - one constant-time check per problem.
- Space: `O(1)` - only a few integer variables are used.

### Implementations
- [Rust solution](src/main.rs)
