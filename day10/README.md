# Day 10: Hoof It

[Puzzle Description](https://adventofcode.com/2024/day/10)

## Problem Description

### Part 1

You discover a reindeer at a Lava Production Facility holding a "Lava Island
Hiking Guide". The reindeer gives you a blank topographic map and you need to
help fill in the hiking trails. The map shows heights from 0 (lowest) to 9
(highest).

A good hiking trail:

- Starts at height 0 (trailhead)
- Ends at height 9
- Always increases by exactly 1 height at each step
- Only moves up, down, left, or right (no diagonals)

A trailhead's score is the number of 9-height positions reachable from that
trailhead via hiking trails.

### Part 2

The second part introduces a new measurement called a trailhead's rating. A
trailhead's rating is the number of distinct hiking trails that begin at that
position. Each trail must follow the same rules as Part 1, but now we need to
count all possible unique paths from each trailhead.

## Solution Notes

The solution uses TypeScript (Deno) and implements:

- Path finding algorithms to discover valid hiking trails
- Counting mechanisms for both trailhead scores and ratings
- Graph traversal to find all possible paths from each trailhead

### Implementation Details

- `mod.ts`: Main solution module
- `mod_test.ts`: Test cases
- `input.txt`: Puzzle input

## Running the Solution

```bash
deno test -A --watch
```

## Results

- Part 1: 482 (sum of all trailhead scores)
- Part 2: 1094 (sum of all trailhead ratings)
