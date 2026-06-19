# 0628 — Log a string field

Using Go's stdlib structured-logging library `log/slog`, this lesson emits an INFO record `login` with one structured string field `user=alice`. The record is captured IN-MEMORY by pointing a `slog.NewJSONHandler` at a `bytes.Buffer` with a `ReplaceAttr` hook that strips the `time` key (no real timestamp). The captured JSON line is then parsed with `encoding/json`, the level is normalized to the short lowercase set, and the fields are sorted by key to print the normalized line: `info|login|user=alice`.

## Run

    go run .
