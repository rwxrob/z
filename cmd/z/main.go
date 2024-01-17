package main

import (
	"os"

	"github.com/rwxrob/z"
)

func main() {
	err := z.Cmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
