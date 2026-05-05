## 116A - Tram

### Description
A tram starts empty and passes through `n` stops. At each stop, some passengers leave and some enter.

The goal is to find the maximum number of passengers inside the tram at any point during the route.

### Mathematical Logic
Keep a running count of passengers currently in the tram:
- subtract the number of passengers who leave
- add the number of passengers who enter

After each stop, update the maximum value seen so far.

### Complexity
- Time: `O(n)` - each stop is processed once.
- Space: `O(1)`.

### Implementations
- [Rust solution](src/main.rs)
