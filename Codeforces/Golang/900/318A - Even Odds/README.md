## 318A - Even Odds

### Description
Given n and k, find the k-th positive integer in the sequence of positive integers where all odd numbers come first (1,3,5,...), followed by even numbers (2,4,6,...).

### Mathematical Logic
- Count how many odd numbers exist among first n: `odd = (n+1)/2`.
- If `k <= odd` the answer is `2*k - 1`, otherwise `2*(k - odd)`.

### Complexity
- Time: `O(1)`.
- Space: `O(1)`.

### Implementations
- [Go solution](main.go)