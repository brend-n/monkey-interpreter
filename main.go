package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/brend-n/monkey-interpreter/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the monkey programming language!\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
