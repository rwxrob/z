package main

import (
	"strings"

	"github.com/rwxrob/bonzai/help"
	Z "github.com/rwxrob/bonzai/z"
)

var tmux = &Z.Cmd{
	Name:     `tmux`,
	Summary:  `make tmux updates`,
	Commands: []*Z.Cmd{help.Cmd, tmuxUpdate},
}

var tmuxUpdate = &Z.Cmd{
	Name:    `update`,
	Summary: `update the onscreen status`,
	Call: func(_ *Z.Cmd, args ...string) error {
		msg := strings.Join(args, " ")
		return Z.Exec(
			"tmux", "-L", "live", "set", "-g", "status-right", msg,
		)
	},
}
