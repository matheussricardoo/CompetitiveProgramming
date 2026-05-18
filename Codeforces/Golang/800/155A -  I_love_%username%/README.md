# 155A - I_love_%username%

## Problem Description

Given a sequence of `n` integer scores, count how many times a new maximum or new minimum value appears in the sequence. The first score is always considered as both the initial maximum and minimum.

## Mathematical Logic

The algorithm maintains running maximum and minimum values:
1. Initialize both max and min to the first score
2. For each subsequent score:
   - If it's greater than current max: increment counter, update max
   - Else if it's less than current min: increment counter, update min
   - Otherwise: no change to counter

The key insight is that each new extreme (maximum or minimum) is counted exactly once when it first appears. Since we only count when a score strictly exceeds or falls below current extremes, ties are ignored.

## Complexity Analysis

- **Time Complexity**: O(n) - single pass through the sequence
- **Space Complexity**: O(n) - for storing the scores array

## Implementation

- [Go solution](main.go)
