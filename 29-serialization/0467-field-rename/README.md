# 0467 — Field rename

Uses the standard library `encoding/json` package with a struct tag to map a
Go field name to a different JSON key. The `Person.FullName` field carries the
tag `json:"full_name"`, so `json.Marshal` emits the key `full_name` instead of
the default field name. This is `encoding/json`'s rename feature, analogous to
serde `rename`, Jackson `@JsonProperty`, or Swift `CodingKeys`.

## Run

    go run .
