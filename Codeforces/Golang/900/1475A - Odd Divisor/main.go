package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	t := readInt(reader)

	for i := 0; i < t; i++ {
		n := readInt(reader)

		if n&(n-1) != 0 {
			fmt.Fprintln(writer, "YES")
		} else {
			fmt.Fprintln(writer, "NO")
		}
	}
}

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
