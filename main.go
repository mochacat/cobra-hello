package main

import (
    "log"

    "github.com/mochacat/cobra-hello/cmd"
)

func main(){
    if err := cmd.RootCmd().Execute(); err != nil {
        log.Fatal(err)
    }
}
