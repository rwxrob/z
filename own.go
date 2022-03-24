package foo

// Go treats all files as if they are, more or less, in the same large
// file. Create separate files to help you and others find the code you
// need quickly.

import (
	"log"

	// if typing bonzai becomes tedious use capital Z as a convention

	Z "github.com/rwxrob/bonzai"
)

var own = &Z.Cmd{
	Name: `own`,
	Call: func(caller *Z.Cmd, none ...string) error {
		log.Print("I'm in my own file.")
		return nil
	},
}
