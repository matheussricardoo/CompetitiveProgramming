## 151A - Soft Drinking

### Description
Given the amount of ingredients available, determine how many toasts each friend can make for a party.

For each toast, a certain amount of drink, lemon slices, and salt is needed. The goal is to maximize the number of toasts per friend.

### Mathematical Logic
Compute how many toasts can be made from each ingredient:
- drink: `(k * l) / nl`
- lemon slices: `c * d`
- salt: `p / np`

The total number of toasts is the minimum of these three values.
Then divide by `n` to get the number of toasts per friend.

### Complexity
- Time: `O(1)` - only arithmetic operations.
- Space: `O(1)`.

### Implementations
- [Rust solution](src/main.rs)
