package main

import (
	"chijason99/northwet_interpreter/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the NorthWet programming language!\n", user.Username)

	repl.Start(os.Stdin, os.Stdout)
}