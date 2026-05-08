## 996A - Hit the Lottery

### Description
Given an amount of money `n`, find the minimum number of bills needed to make that exact amount using denominations: `100`, `20`, `10`, `5`, `1`.

### Mathematical Logic
Use a greedy approach: process bills in descending order. For each denomination, use as many bills as possible, then move to the next smaller denomination.

### Complexity
- Time: `O(1)` - constant number of denominations.
- Space: `O(1)`.

### Implementations
- [Go solution](main.go)
## 996A - Hit the Lottery

### Description
Given an integer `n`, find the minimum number of bills needed to make amount `n` using denominations of `100`, `20`, `10`, `5`, and `1`.

### Mathematical Logic
Use a greedy approach: repeatedly use the largest bill that doesn't exceed the remaining amount.

For each bill (in descending order):
- divide `n` by the bill to get the count
- subtract using modulo to find the remainder

### Complexity
- Time: `O(1)` - fixed number of denominations.
- Space: `O(1)`.

### Implementations
- [Go solution](main.go)
