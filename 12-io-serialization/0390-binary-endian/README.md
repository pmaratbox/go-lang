# 0390 — Big-Endian Bytes

Encode the integer 258 as two big-endian bytes (1, 2), decode them back to 258, printing `1 2 258`. Go's encoding/binary.BigEndian handles uint16 round-trips directly.

## Run

    go run .
