package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	current, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! this os the Monkey programming lengauge!\n", current.Username)
	fmt.Printf("Fell free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
