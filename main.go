package main

import (
	"fmt"
	"gox509inspector/inspector"
	"os"
)

func main() {
	// pipe certificate in
	c, err := inspector.Parsex509Cert(os.Stdin)
	if c == nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
	}

	bytes := inspector.PrintX509Cert(c)
	fmt.Print(string(bytes))
}
