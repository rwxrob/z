package main

import (
	"log"

	"github.com/rwxrob/bonzai"
	Z "github.com/rwxrob/bonzai"
	"github.com/rwxrob/bonzai/inc/help"
	"github.com/rwxrob/config"
	"github.com/rwxrob/twitch"
	"github.com/rwxrob/y2j"
	"github.com/rwxrob/yq"
)

func main() {
	log.SetFlags(0)
	bonzai.Aliases = map[string][]string{
		"status":   {"tmux", "update"},
		"project":  {"twitch", "bot", "commands", "edit", "project"},
		"commands": {"twitch", "bot", "commands"},
		"work":     {"go", "work"},
		"chat":     {"twitch", "chat"},
	}
	Cmd.Run()
}

var Cmd = &Z.Cmd{
	Name:      `z`,
	Summary:   `rwxrob's bonzai command tree`,
	Version:   `v0.0.1`,
	Copyright: `Copyright 2021 Robert S Muhlestein`,
	License:   `Apache-2.0`,
	Commands: []*Z.Cmd{
		help.Cmd, config.Cmd, y2j.Cmd, twitch.Cmd, tmux, yq.Cmd, golang,
	},
}
