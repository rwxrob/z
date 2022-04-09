package main

import (
	"log"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/config"
	"github.com/rwxrob/help"
	"github.com/rwxrob/twitch"
	"github.com/rwxrob/uniq"

	//"github.com/rwxrob/update"
	"github.com/rwxrob/y2j"
	"github.com/rwxrob/yq"
)

func main() {

	// remove log prefixes
	log.SetFlags(0)

	// provide panic trace
	Z.AllowPanic = true

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
		"commands": {"twitch", "bot", "commands", "file", "edit"},
		"sync":     {"twitch", "bot", "commands", "sync"},
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
	Version:   `v0.1.5`,
	Copyright: `Copyright 2021 Robert S Muhlestein`,
	License:   `Apache-2.0`,
	Commands: []*Z.Cmd{
		help.Cmd, config.Cmd, y2j.Cmd, twitch.Cmd, tmux, yq.Cmd, golang,
		uniq.Cmd, //update.Cmd,
	},
	Description: `
		Hi, I'm rwxrob <http://rwxrob.tv> and this is my Bonzai™ tree. I am
		slowly replacing all my shell scripts and other Go utilities with
		Bonzai branches that I graft into this *z* command. You are welcome
		to play around with it, but please know that I am radically changing
		things *daily*.`,
}
