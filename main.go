package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! this os the Monkey programming lengauge!\n", user.Username)
	fmt.Printf("Fell free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
