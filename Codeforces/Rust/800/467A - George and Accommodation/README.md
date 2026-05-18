## 467A - George and Accommodation

### Description
George has `n` friends who want to stay with him. Each friend requests accommodation in a room, and George needs to check if the room can accommodate the friend.

For each friend, the room has size `p` and the friend needs space `q`. The friend can stay if `|p - q| >= 2`.

Count the number of friends that can be accommodated.

### Mathematical Logic
For each room `i` with capacity `p[i]` and friend `i` needing space `q[i]`:
- if `|p[i] - q[i]| >= 2`, the friend can stay (increment the count)

Total count = number of accommodating rooms.

### Complexity
- Time: `O(n)` - each room is checked once.
- Space: `O(1)`.

### Implementations
- [Rust solution](src/main.rs)
