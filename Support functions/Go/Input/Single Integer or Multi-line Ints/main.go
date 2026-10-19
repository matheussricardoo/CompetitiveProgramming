// Minimal integer reader example; EOF returns an error instead of looping forever.
package main

import (
    "bufio"
    "fmt"
    "os"
)

func readInt(reader *bufio.Reader) (int, error) {
    var value int
    _, err := fmt.Fscan(reader, &value)
    return value, err
}

func main() {
    if value, err := readInt(bufio.NewReader(os.Stdin)); err == nil {
        fmt.Println(value)
    }
}
