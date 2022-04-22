package main

import (
	"log"
	"text/template"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/conf"
	"github.com/rwxrob/help"
	"github.com/rwxrob/twitch"
	"github.com/rwxrob/uniq"
	"github.com/rwxrob/vars"
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
			// "conf": {conf.Cmd}, // bork cuz no multicall mode
			"yq":     {yq.Cmd},
			"y2j":    {y2j.Cmd},
			"status": {tmux, "update"},
			"afk":    {twitch.Cmd, "chat", "!afk"},
		}
		Z.Run()

	*/

	// MONOLITH (z) - which I prefer
	Z.Shortcuts = map[string][]string{
		"status":    {"tmux", "update"},
		"offscreen": {"chat", "!offscreen"},
		"project":   {"twitch", "bot", "commands", "edit", "project"},
		"commands":  {"twitch", "bot", "commands", "file", "edit"},
		"sync":      {"twitch", "bot", "commands", "sync"},
		"work":      {"go", "work"},
		"chat":      {"twitch", "chat"},
		"afk":       {"twitch", "chat", "!afk"},
		"isosec":    {"uniq", "isosec"},
		"isonan":    {"uniq", "isonan"},
		"uuid":      {"uniq", "uuid"},
		"epoch":     {"uniq", "second"},
	}

	Cmd.Run()

}

var Cmd = &Z.Cmd{
	Name:      `z`,
	Summary:   `rwxrob's bonzai command tree`,
	Version:   `v0.2.2`,
	Copyright: `Copyright 2021 Robert S Muhlestein`,
	License:   `Apache-2.0`,
	Site:      `rwxrob.tv`,
	Source:    `git@github.com:rwxrob/z.git`,
	Issues:    `github.com/rwxrob/z/issues`,

	Commands: []*Z.Cmd{
		help.Cmd, conf.Cmd, vars.Cmd, y2j.Cmd, twitch.Cmd, tmux, yq.Cmd, golang,
		uniq.Cmd, //github.Cmd, //update.Cmd,
	},

	Dynamic: template.FuncMap{
		"uname": func() string { return Z.Out("uname", "-a") },
		"ls":    func() string { return Z.Out("ls", "-l", "-h") },
	},

	Description: `
		Hi, I'm rwxrob.tv and this is my Bonzai™ tree. I am slowly
		replacing all my shell scripts and other Go utilities with Bonzai
		branches that I graft into this {{cmd .Name}} command. You are welcome
		to play around with it, but please know that I am radically changing
		things *daily*.

		Also check out https://github.com/rwxrob/foo for a sample template
		Bonzai tree to get started on your own.

		Here's some random output from the Dynamic *ls* function piped to
		the builtin *indent* function using {{ "{{ ls | indent 4 }}" }} Go
		template syntax:

		{{ ls | indent 4 }}
		
		That was a verbatim block because of the indent.

		`,
}
