## 271A - Beautiful Year

### Description
Given a year `y`, find the next year that has all different digits in its decimal representation.

A year is "beautiful" if all its digits are distinct (no repeated digits).

### Mathematical Logic
Iterate from `y + 1` upwards. For each year, extract the four digits:
- `a = year / 1000`
- `b = (year / 100) % 10`
- `c = (year / 10) % 10`
- `d = year % 10`

Check if all are distinct: `a != b && a != c && a != d && b != c && b != d && c != d`.

Return the first year that satisfies this condition.

### Complexity
- Time: `O(k)` - where `k` is the number of years to search (typically small).
- Space: `O(1)`.

### Implementations
- [Go solution](main.go)
