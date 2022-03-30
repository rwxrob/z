package main

import (
	"log"

	Z "github.com/rwxrob/bonzai"
	"github.com/rwxrob/bonzai/inc/help"
	"github.com/rwxrob/config"
	"github.com/rwxrob/twitch"
	"github.com/rwxrob/uniq"
	"github.com/rwxrob/y2j"
	"github.com/rwxrob/yq"
)

func main() {
	log.SetFlags(0)

	// can run in multicall, or monolith, not both

	/*

		// MULTICALL (status, afk, etc. linked)
		// (no completion unless set for individual commands)
		// (requires creation of hard/sym links or copies)
		Z.Commands = map[string][]any{
			// "config": {config.Cmd}, // bork cuz no multicall mode
			"yq":     {yq.Cmd},
			"y2j":    {y2j.Cmd},
			"status": {tmux, "update"},
			"afk":    {twitch.Cmd, "chat", "!afk"},
		}
		Z.Run()

	*/

	// MONOLITH (z) - which I prefer
	Z.Aliases = map[string][]string{
		"status":   {"tmux", "update"},
		"project":  {"twitch", "bot", "commands", "edit", "project"},
		"commands": {"twitch", "bot", "commands"},
		"work":     {"go", "work"},
		"chat":     {"twitch", "chat"},
		"afk":      {"twitch", "chat", "!afk"},
		"isosec":   {"uniq", "isosec"},
		"uuid":     {"uniq", "uuid"},
		"epoch":    {"uniq", "second"},
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
		help.Cmd, config.Cmd, y2j.Cmd, twitch.Cmd, tmux, yq.Cmd, golang, uniq.Cmd,
	},
}
