package main

import (
	"fmt"
	"os"

	"github.com/indiependente/gox509inspector/inspector"
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

	bytes := inspector.GetQuickInfo(c)
	fmt.Print(string(bytes))
}
