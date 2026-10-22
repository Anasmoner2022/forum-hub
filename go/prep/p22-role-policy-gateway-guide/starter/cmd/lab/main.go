package main

import (
	"fmt"
	"os"
	"academy.local/p22/internal/lab"
)

func main() {
	if err := lab.Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
