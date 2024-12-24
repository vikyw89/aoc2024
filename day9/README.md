# Day 9: Disk Defragmentation

## Overview

This document provides a detailed explanation of the solution for Day 9 of the
Advent of Code 2024 challenge. This challenge involves simulating disk
defragmentation based on a compressed disk map.

## Problem Description

The challenge involves processing a compressed disk map represented as a string
of alternating file block sizes and space sizes. The goal is to decompress this
map into an array of file IDs and spaces (-1), then perform a series of
operations to simulate defragmentation.

**Part 1:** The first part requires decompressing the disk map, calculating a
checksum based on the decompressed map, and outputting the checksum. The
checksum is calculated by multiplying each file ID by its index and summing the
results, ignoring spaces (-1).

**Part 2:** The second part builds upon the first. After decompressing the disk
map, the challenge requires moving the last file block to the front of the disk
map, effectively defragmenting the disk. This process involves finding the last
file block, locating a space large enough to accommodate it, and swapping the
file block with the space. This process is repeated for each file block, moving
from the end of the disk map to the beginning. Finally, a checksum is calculated
on the defragmented disk map.

## Solution

### Part 1

- The `decompressDiskmap` function takes the compressed disk map string and
  returns a decompressed array of file IDs and spaces (-1).
- The `getCheckSum` function calculates the checksum of the decompressed array
  by multiplying each file ID by its index and summing the results, ignoring
  spaces (-1).

### Part 2

- The `rollFileBlockToFront` function implements the defragmentation logic. It
  iterates through the decompressed disk map from the end to the beginning. For
  each file block, it finds the first available space large enough to
  accommodate the file block and swaps the file block with the space.
- The `getSpaceIndex` and `getFileBlock` helper functions are used to locate the
  space and file block respectively.
- Finally, the `getCheckSum` function is used again to calculate the checksum of
  the defragmented disk map.

## How to Run

To execute the solution for Day 9, navigate to the `day9` directory and run the
following command:

```bash
deno test -A --watch
```
