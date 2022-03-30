package main

import (
	"log"
	"os"

	Z "github.com/rwxrob/bonzai"
	"github.com/rwxrob/bonzai/inc/help"
	"github.com/rwxrob/fs/file"
)

var golang = &Z.Cmd{
	Name:     `go`,
	Summary:  `go related helper actions`,
	MinArgs:  1,
	Commands: []*Z.Cmd{help.Cmd, gowork},
}

var gowork = &Z.Cmd{
	Name:     `work`,
	Summary:  `turn on or off go.work file`,
	Usage:    `(on|off)`,
	MinArgs:  1,
	Params:   []string{"on", "off"},
	Commands: []*Z.Cmd{help.Cmd},
	Call: func(x *Z.Cmd, args ...string) error {
		switch args[0] {
		case "on":
			if file.Exists("go.work.off") {
				log.Print("go.work.off -> go.work")
				return os.Rename("go.work.off", "go.work")
			}
		case "off":
			if file.Exists("go.work") {
				log.Print("go.work -> go.work.off")
				return os.Rename("go.work", "go.work.off")
			}
		default:
			return x.UsageError()
		}
		return nil
	},
}
