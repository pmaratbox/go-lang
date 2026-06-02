# 0080 — Deduplicate

Remove duplicates from `1, 2, 2, 3, 1`, keeping the first occurrence of each in order, and print `1 2 3`. A `seen` map tracks values already added; each new value is appended once, preserving first-seen order.

## Run

    go run .
