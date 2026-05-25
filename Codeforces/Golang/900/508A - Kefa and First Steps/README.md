## 508A - Kefa and First Steps

### Description
Given a sequence of integers, find the length of the longest non-decreasing contiguous subsegment.

### Mathematical Logic
Traverse the array while maintaining the current length of the non-decreasing streak and the maximum length seen so far. If the current value is at least the previous one, extend the streak; otherwise, reset it.

### Complexity
- Time: `O(n)` - one pass through the array.
- Space: `O(1)`.

### Implementations
- [Go solution](main.go)
