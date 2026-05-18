## 344A - Magnets

### Description
A sequence of magnets is arranged in a line, each with a polarity (either `01` or `10`).

The goal is to count how many groups of consecutive magnets with the same polarity exist.

### Mathematical Logic
Iterate through the magnets and count every transition where the polarity changes from one magnet to the next:

- Initialize `count = 0` and `previous = first_magnet`
- for each subsequent magnet:
  - if magnet polarity != previous, increment count
  - update previous to current polarity
- result = count + 1 (the +1 accounts for the first group)

Alternatively, directly count transitions and add 1.

### Complexity
- Time: `O(n)` - each magnet is processed once.
- Space: `O(1)`.

### Implementations
- [Go solution](main.go)
