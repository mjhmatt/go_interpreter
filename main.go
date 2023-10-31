package main

import (
	"fmt"
	"go_interpreter/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("hello %s! This is Monkeylang", user.Username)
	fmt.Printf("Type a command")
	repl.Start(os.Stdin, os.Stdout)

}
