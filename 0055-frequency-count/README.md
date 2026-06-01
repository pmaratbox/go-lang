# 0055 — Frequency Count

Count how many times each letter appears in `banana` and print the per-letter counts in alphabetical order: `a:3 b:1 n:2`. A `map[rune]int` tallies the counts (`counts[ch]++` zero-initializes a missing key). Map iteration order is randomized, so the keys are collected and `sort`ed before printing.

## Run

    go run .
