## 236A - Boy or Girl

### Description
The goal is to determine the gender based on the number of distinct characters in a name.

- If the number of distinct characters is even, print `CHAT WITH HER!`
- If the number of distinct characters is odd, print `IGNORE HIM!`

### Mathematical Logic
Given a name string, count the number of distinct characters it contains.

Let `distinct_count = |{unique characters in name}|`.

Then:
- if `distinct_count % 2 == 0`, answer is `CHAT WITH HER!`
- otherwise, answer is `IGNORE HIM!`

### Complexity
- Time: `O(n)` - each character is processed once to build the set of distinct characters.
- Space: `O(k)` - where `k` is the number of distinct characters in the name (at most 26 for lowercase letters).

### Implementations
- [Go solution](main.go)
