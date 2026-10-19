// Read one whitespace-delimited token. Empty input terminates normally.
package main

import (
    "bufio"
    "fmt"
    "os"
)

func readString(reader *bufio.Reader) (string, error) {
    var value string
    _, err := fmt.Fscan(reader, &value)
    return value, err
}

func main() {
    if value, err := readString(bufio.NewReader(os.Stdin)); err == nil {
        fmt.Println(value)
    }
}
