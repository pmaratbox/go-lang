# 0228 — Parse Quoted CSV

Parse the CSV row `a,"b,c",d`, respecting the quoted comma, into three fields joined by pipes `a|b,c|d`. A raw string literal holds the input and a bool tracks the in-quotes state.

## Run

    go run .
