# 0423 — EventEmitter (Pub/Sub)

Build a multi-topic EventEmitter with on(topic, handler), emit(topic, payload), and off(topic, handler). A map of topic to handler slices, with off filtering in place via a zero-alloc `kept := hs[:0]` reslice.

## Run

    go run .
