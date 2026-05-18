## 1676A - Lucky?

### Description
A ticket number (6 digits) is considered "lucky" if the sum of the first three digits equals the sum of the last three digits.

Given a 6-digit string representing a ticket number, determine if it's lucky.

### Mathematical Logic
1. Extract the 6 digits from the ticket string.
2. Calculate the sum of the first three digits: `sum_first = d[0] + d[1] + d[2]`
3. Calculate the sum of the last three digits: `sum_last = d[3] + d[4] + d[5]`
4. If `sum_first == sum_last`, the ticket is lucky (output "YES"), otherwise output "NO".

### Complexity
- Time: `O(1)` - constant number of operations (exactly 6 digits).
- Space: `O(1)`.

### Implementations
- [Rust solution](src/main.rs)
