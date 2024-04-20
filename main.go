package main

import (
	"fmt"
	"os"
	"os/user"
	"pschwarz/interpreter/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey Programming Language\n",
		user.Username)

	fmt.Printf("Have fun trying out the language by typing below\n")
	repl.Start(os.Stdin, os.Stdout)
}
