package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/rwxrob/bonzai/help"
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/fs/file"
)

var golang = &Z.Cmd{
	Name:     `go`,
	Summary:  `go related helper actions`,
	Aliases:  []string{"rust"},
	MinArgs:  1,
	Commands: []*Z.Cmd{help.Cmd, gowork, genisrune},
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

var genisrune = &Z.Cmd{
	Name:     `genisrune`,
	Summary:  `generate a performant rune check condition`,
	Commands: []*Z.Cmd{help.Cmd},
	MinArgs:  1,
	Call: func(x *Z.Cmd, args ...string) error {
		var conds []string
		for _, r := range args[0] {
			conds = append(conds, fmt.Sprintf("r==%q", r))
		}
		fmt.Println(strings.Join(conds, "||"))
		return nil
	},
}
