package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	var buf [2]byte
	binary.BigEndian.PutUint16(buf[:], 258)

	decoded := binary.BigEndian.Uint16(buf[:])

	fmt.Printf("%d %d %d\n", buf[0], buf[1], decoded)
}
