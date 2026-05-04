## 734A - Anton and Danik

### Description
Anton and Danik play a game. A string of length `n` consisting of `A` and `D` characters is given.

Count the occurrences of `A` and `D`:
- if `A` count > `D` count, print `Anton`
- if `D` count > `A` count, print `Danik`
- if both counts are equal, print `Friendship`

### Mathematical Logic
Let:
- `count_a` = number of `A` characters
- `count_d` = number of `D` characters

Then:
- if `count_a > count_d`, winner is `Anton`
- if `count_d > count_a`, winner is `Danik`
- otherwise, the result is `Friendship`

### Complexity
- Time: `O(n)` — each character is counted once.
- Space: `O(1)` — only two counters are used.

### Implementations
- [Go solution](main.go)
