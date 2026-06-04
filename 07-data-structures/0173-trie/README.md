# 0173 — Trie

Insert "cat" and "car" into a trie, then search "car" (yes) and "can" (no), printing `yes no`. Go models each node with a `map[rune]*trieNode` of children plus an end-of-word flag.

## Run

    go run .
