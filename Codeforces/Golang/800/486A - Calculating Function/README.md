## 486A - Calculating Function

### Description
Given an integer `n`, compute the value of the alternating sum:

- `f(n) = 1 - 2 + 3 - 4 + 5 - 6 + ... +/- n`

### Mathematical Logic
Observe the pattern:
- if `n` is even: pairs `(1-2) + (3-4) + ... + ((n-1) - n) = -1 - 1 - ... - 1 = -n/2`, so the result is `n/2`.
- if `n` is odd: the sum is `1 - 2 + 3 - 4 + ... + n = 1 + (-2 + 3) + (-4 + 5) + ... + (-((n-1) + n) = 1 + 1 + 1 + ... = -(n+1)/2`.

Result:
- even: `n / 2`
- odd: `-(n + 1) / 2`

### Complexity
- Time: `O(1)` - constant arithmetic.
- Space: `O(1)`.

### Implementations
- [Go solution](main.go)
