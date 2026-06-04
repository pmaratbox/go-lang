# 0164 — Merge Maps

Merge {a:1,b:2} and {b:3,c:4} with the right map winning on conflicts, printing `a:1 b:3 c:4`. In Go you copy the left map then overwrite from the right, since map iteration order is randomized you sort the keys before printing.

## Run

    go run .
