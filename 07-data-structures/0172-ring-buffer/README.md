# 0172 — Ring Buffer

Push 1,2,3,4,5 into a fixed capacity-3 ring buffer (overwriting oldest) and print the final contents `3 4 5`. Go backs the buffer with a fixed-length slice and modular head/count indices.

## Run

    go run .
