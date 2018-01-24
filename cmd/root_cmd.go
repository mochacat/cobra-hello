package cmd

import (
    "fmt"
    "os"
    "io"

    "github.com/spf13/cobra"
    "github.com/spf13/viper"
)

var out io.Writer = os.Stdout

type name struct {
    value string
}

func (n name) printName(){
    fmt.Fprintf(out, "Hello, %#s!\n", n.value)
}

func sayHello(cmd *cobra.Command, args []string) {
    n := name{viper.GetString("name")}
    n.printName()
}

func RootCmd() *cobra.Command {
    rootCmd := cobra.Command{
        Use: "hello",
        Short: "greet a person",
    }

    rootCmd.Run = sayHello
    rootCmd.Flags().String("name", "", "Oh hey it is a name")
    viper.BindPFlag("name", rootCmd.Flags().Lookup("name"))
    return &rootCmd
}
