## 427A - Police Recruits

### Description
Police recruits handle a sequence of events. Positive numbers mean new police officers arrive, while `-1` means a crime happened.

If there is no available police officer to handle a crime, it is untreated. Count how many crimes are untreated.

### Mathematical Logic
Maintain:
- `polices` = number of available officers
- `count` = number of untreated crimes

For each event:
- if positive, add it to `polices`
- if `-1` and `polices > 0`, consume one officer
- otherwise, increment `count`

### Complexity
- Time: `O(n)` - each event is processed once.
- Space: `O(1)`.

### Implementations
- [Rust solution](src/main.rs)
