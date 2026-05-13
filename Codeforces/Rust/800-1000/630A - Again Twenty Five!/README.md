# 630A - Again Twenty Five!

## Problem Description

Read an integer `n` and output a value related to number theory properties of integers. Despite taking an integer input, the output is always a fixed value.

## Mathematical Logic

This problem is based on the mathematical property that any number ending in 5, when squared, always ends in 25. Since the problem provides only one input, the answer is always constant. Specifically, n² mod 100 always equals 25 when n ≡ 5 (mod 10), which is the constraint implied by the problem.

The solution simply outputs the fixed value without complex computation, demonstrating the elegance of mathematical shortcuts in competitive programming.

## Complexity Analysis

- **Time Complexity**: O(1) - constant time operation (no computation needed)
- **Space Complexity**: O(1) - no extra space required

## Implementation

- [Rust solution](src/main.rs)
