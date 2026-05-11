## 723A - The New Year: Meeting Friends

### Description
Given three positions of friends on a line, find the minimum distance they all need to move so they can meet at the same point.

### Mathematical Logic
Sort the three positions and compute the distance between the maximum and the minimum values.

That distance is the minimum total movement needed.

### Complexity
- Time: `O(1)` - sorting three numbers is constant work.
- Space: `O(1)`.

### Implementations
- [Go solution](main.go)
