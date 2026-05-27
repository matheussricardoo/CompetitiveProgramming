package main

import (
	"bufio"
)

func readLine(reader *bufio.Reader) string {
	line, _, _ := reader.ReadLine()
	return string(line)
}
