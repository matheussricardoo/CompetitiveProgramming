## 110A - Nearly Lucky Number

### Description
A number is considered "lucky" if its decimal representation contains only the digits `4` and `7`.

Given a number `n`, count how many times the digits `4` and `7` appear in it.
If this count is itself `4` or `7`, print `YES`. Otherwise, print `NO`.

### Mathematical Logic
Let `lucky_count` = number of `4`s and `7`s in the decimal representation of `n`.

Then:
- if `lucky_count == 4` or `lucky_count == 7`, print `YES`
- otherwise, print `NO`

### Complexity
- Time: `O(d)` — where `d` is the number of digits in `n`.
- Space: `O(1)` — only a counter is used.

### Implementations
- [Rust solution](src/main.rs)
