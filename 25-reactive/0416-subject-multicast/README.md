# 0416 — Subject Multicast

Implement a Subject that multicasts each emission to all current observers; two observers both receive 1 then 2. The Subject holds a slice of observer closures and fans each value out in subscription order.

## Run

    go run .
