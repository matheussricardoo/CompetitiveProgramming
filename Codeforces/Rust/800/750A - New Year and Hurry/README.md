# 750A - New Year and Hurry

## Problem Description

Given `n` problems with increasing difficulty (problem `i` takes `5 * i` minutes to solve) and `k` minutes already spent, determine how many problems can be solved with the remaining time. You have a total of 240 minutes available.

## Mathematical Logic

The algorithm iterates through problems 1 to n sequentially. For each problem i:
- Calculate the time required: `5 * i` minutes
- Check if `timeRemaining >= 5 * i`
- If yes, solve it and decrement `timeRemaining` by `5 * i`
- If no, stop (cannot solve remaining problems due to insufficient time)

The key insight is that problems must be solved in order, so as soon as we cannot solve problem i, we cannot solve any problem j > i either (since they take more time).

## Complexity Analysis

- **Time Complexity**: O(n) - iterate through at most n problems
- **Space Complexity**: O(1) - only use constant extra space for counters

## Implementation

- [Rust solution](src/main.rs)
