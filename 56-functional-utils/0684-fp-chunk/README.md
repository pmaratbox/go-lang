# 0684 — Chunk

Uses the [samber/lo](https://github.com/samber/lo) functional library and its `lo.Chunk` transform to split `[1, 2, 3, 4, 5, 6]` into fixed-size pieces of 2. Each chunk is comma-joined, and the chunks are joined with `|`.

## Run

    go run .
