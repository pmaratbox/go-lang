# 0323 — Maybe Monad

Chain Maybe operations: Some(2) then +3 then *2 gives 10, and a None chain yields the fallback, printing `10 none`. A small struct with a `present` flag plays the part of an optional, and `bind` short-circuits on None.

## Run

    go run .
