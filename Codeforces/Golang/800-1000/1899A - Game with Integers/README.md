# 1899A - Game with Integers

## Problem Description

Two players play a game on an integer `n` where each can either add or subtract 1. The first player who cannot make a move (when reaching a losing state) loses. Determine if the first player wins based on the divisibility of `n` by 3.

## Mathematical Logic

The key insight is that numbers not divisible by 3 are winning positions for the first player, while numbers divisible by 3 are losing positions for the first player.

This is because:
- If n is not divisible by 3, the first player can make a move to leave the second player with a multiple of 3
- If n is divisible by 3, any move (±1) results in a non-multiple of 3, giving the second player a winning position

Therefore: First player wins if n % 3 ≠ 0, otherwise second player wins.

## Complexity Analysis

- **Time Complexity**: O(t) where t is the number of test cases (each test is O(1))
- **Space Complexity**: O(1) - only constant space for variables

## Implementation

- [Go solution](main.go)
