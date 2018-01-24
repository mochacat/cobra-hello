package cmd

import (
    "bytes"
    "testing"
)

func TestSayHello(t *testing.T) {
    cases := []struct {
        name, want string
    }{
        {"Cat", "Hello, Cat!\n"},
        {"", "Hello, !\n"},
        {"Mr. Fancy Pants", "Hello, Mr. Fancy Pants!\n"},
    }
    for _, c := range cases {
        bak := out
        out = new(bytes.Buffer)
        defer func() { out = bak }()
        n := name{c.name}
        n.printName()
        printedName := out.(*bytes.Buffer).String()
        if printedName != c.want {
            t.Errorf("Wanted %q, Got %q", c.want, printedName)
        }
    }
}
