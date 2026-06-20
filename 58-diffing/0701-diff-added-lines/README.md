# 0701 — Added line content

Using the `pmezard/go-difflib` library's `SequenceMatcher`, this program diffs
two fixed line-lists and walks the resulting opcodes. Lines that the matcher
tags as inserted (`'i'`) or as the target side of a replacement (`'r'`) are the
ADDED lines; they are collected in document (B) order and printed comma-joined.

## Run

    go run .
