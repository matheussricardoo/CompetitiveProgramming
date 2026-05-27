package main

import (
	"bufio"
)

func readString(reader *bufio.Reader) string {
	b, _ := reader.ReadByte()
	for b <= ' ' {
		b, _ = reader.ReadByte()
	}
	var res []byte
	for b > ' ' {
		res = append(res, b)
		b, _ = reader.ReadByte()
	}
	return string(res)
}