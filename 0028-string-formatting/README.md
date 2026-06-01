# 0028 — String Formatting

Format the float `3.14159` to two decimals and zero-pad the integer `42` to width five, printing `pi: 3.14` and `id: 00042`. `fmt.Printf` uses C-style verbs: `%.2f` fixes two decimals and `%05d` zero-pads to width 5. Go's formatting is locale-independent, always using a dot for the decimal point.

## Run

    go run .
