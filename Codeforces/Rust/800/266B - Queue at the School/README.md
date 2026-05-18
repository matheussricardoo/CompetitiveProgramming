## 266B - Queue at the School

### Description
A queue of boys (`B`) and girls (`G`) stands in line. During each second, every adjacent pair `BG` swaps positions simultaneously.

The goal is to simulate this process for `t` seconds and print the final queue.

### Mathematical Logic
For each second:
- scan the queue from left to right
- if a `B` is immediately followed by a `G`, swap them
- skip the next position after a swap to avoid double processing

Repeat this for `t` iterations.

### Complexity
- Time: `O(t * n)` - the queue is scanned `t` times.
- Space: `O(n)` - to store the queue.

### Implementations
- [Rust solution](src/main.rs)
