## 732A - But a Shovel

### Description
You need to buy shovels. Each shovel costs `k` rubles. You have a discount rule: the price ends with digit `r`.

Find the minimum number of shovels you can buy such that the total price (shovels × k) ends with digit `0` or `r`.

### Mathematical Logic
Iterate from 1 to 10 shovels:
1. Calculate total price: `total_price = shovels * k`
2. Check if the last digit of total price is 0 or r: `total_price % 10 == 0 || total_price % 10 == r`
3. Return the first (minimum) number of shovels that satisfies this condition.

Since at most 10 shovels need to be checked, we're guaranteed to find a valid answer within the range [1, 10].

### Complexity
- Time: `O(1)` - at most 10 iterations.
- Space: `O(1)`.

### Implementations
- [Go solution](main.go)
