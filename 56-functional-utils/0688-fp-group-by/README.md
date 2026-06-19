# 0688 — Group by

This lesson uses the `github.com/samber/lo` functional library and its `lo.GroupBy` transform to group `[1,2,3,4,5,6]` by parity into `even`/`odd` buckets. We collect the bucket keys with `lo.Keys`, sort them, then render each as `key:v1,v2,...` joined by `;`.

## Run

    go run .
