## 1352A - Sum of Round Numbers

### Description
Given a number `n`, decompose it into a sum of round numbers.

A round number has exactly one non-zero digit.

### Mathematical Logic
Process each digit of `n` from least significant to most significant:
- if the digit is non-zero, multiply it by the current place value and store it
- move to the next digit

The stored values form the round numbers.

### Complexity
- Time: `O(d)` - where `d` is the number of digits in `n`.
- Space: `O(d)` - to store the resulting round numbers.

### Implementations
- [Rust solution](src/main.rs)
