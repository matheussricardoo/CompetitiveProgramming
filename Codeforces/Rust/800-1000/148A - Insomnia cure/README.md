## 148A - Insomnia cure

### Description
A product is damaged by insects that breed on days that are multiples of `k`, `l`, `m`, or `n`.

Given `d` days, count how many days the product is damaged (days divisible by at least one of the four numbers).

### Mathematical Logic
Iterate through days `1` to `d` and count those where `day % k == 0 || day % l == 0 || day % m == 0 || day % n == 0`.

### Complexity
- Time: `O(d)` - checking each day.
- Space: `O(1)`.

### Implementations
- [Rust solution](src/main.rs)
