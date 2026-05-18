## 266A - Stones on the Table

### Description
Given a sequence of colored stones represented by a string, count the minimum number of stones to remove so that no two adjacent stones have the same color.

### Mathematical Logic
Scan the string and count positions where `s[i] == s[i+1]`. Each such pair requires removing one stone to break the adjacency.

Minimum removals = number of adjacent equal pairs.

### Complexity
- Time: `O(n)` - single pass over the string.
- Space: `O(1)` - constant extra space.

### Implementations
- [Go solution](main.go)
