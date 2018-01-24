package cmd

import (
    "bytes"
    "fmt"
    "testing"
)

func TestSayHello(t *testing.T) {
    bak := out
    out = new(bytes.Buffer)
    defer func() { out = bak }()
    n := name{"Jerk"}
    n.printName()
    if out.(*bytes.Buffer).String() == "Hello, Jerk!\n" {
        fmt.Println("Works as expected")
    }
}
