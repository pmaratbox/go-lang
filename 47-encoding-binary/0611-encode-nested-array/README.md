# 0611 — Encode a nested array

Using the `github.com/vmihailenco/msgpack/v5` MessagePack library, `msgpack.Marshal` encodes the nested array `[[1,2],[3,4]]` to its binary form, which `hex.EncodeToString` renders as lowercase hex: `92920102920304` (an outer fixarray `92` holding two inner fixarrays `92`, each storing its positive fixint elements in a single byte).

## Run

    go run .
