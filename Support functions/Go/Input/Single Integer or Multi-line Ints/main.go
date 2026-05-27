package main

import (
	"bufio"
)

func readInt(reader *bufio.Reader) int {
	b, _ := reader.ReadByte()
	for b <= ' ' {
		b, _ = reader.ReadByte()
	}
	sign := 1
	if b == '-' {
		sign = -1
		b, _ = reader.ReadByte()
	}
	res := 0
	for b >= '0' && b <= '9' {
		res = res*10 + int(b-'0')
		b, _ = reader.ReadByte()
	}
	return res * sign
}