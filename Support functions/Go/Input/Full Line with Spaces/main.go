// Read a complete line, including lines longer than bufio.Reader's buffer.
package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func readLine(reader *bufio.Reader) (string, error) {
    line, err := reader.ReadString('\n')
    return strings.TrimSuffix(strings.TrimSuffix(line, "\n"), "\r"), err
}

func main() {
    line, _ := readLine(bufio.NewReader(os.Stdin))
    if line != "" { fmt.Println(line) }
}
