package main

import (
	"log"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/conf"
	"github.com/rwxrob/github"
	"github.com/rwxrob/good"
	"github.com/rwxrob/help"
	"github.com/rwxrob/keg"
	"github.com/rwxrob/kube"
	"github.com/rwxrob/pomo"
	"github.com/rwxrob/slug"
	"github.com/rwxrob/twitch"
	"github.com/rwxrob/uniq"
	"github.com/rwxrob/vars"
	"github.com/rwxrob/y2j"
	"github.com/rwxrob/yq"
	"github.com/rwxrob/z/env"
	"github.com/rwxrob/z/git"
	"github.com/rwxrob/z/tmux"
)

func init() {
	Z.Dynamic["uname"] = func() string { return Z.Out("uname", "-a") }
	Z.Dynamic["ls"] = func() string { return Z.Out("ls", "-l", "-h") }
}

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

	Cmd.Run()
}

var Cmd = &Z.Cmd{
	Name:      `z`,
	Summary:   `rwxrob's bonzai command tree`,
	Copyright: `Copyright 2021 Robert S Muhlestein`,
	Version:   `v0.1.2`,
	License:   `Apache-2.0`,
	Site:      `rwxrob.tv`,
	Source:    `git@github.com:rwxrob/z.git`,
	Issues:    `github.com/rwxrob/z/issues`,

	Commands: []*Z.Cmd{
		help.Cmd, conf.Cmd, vars.Cmd,
		y2j.Cmd, twitch.Cmd, tmux.Cmd, yq.Cmd,
		uniq.Cmd, pomo.Cmd, github.Cmd, git.Cmd,
		kube.Cmd, env.Cmd, keg.Cmd, slug.Cmd, good.Cmd,
		// openapi.Cmd, update.Cmd, goutil.Cmd
	},

	Shortcuts: Z.ArgMap{
		`project`:   {`twitch`, `bot`, `commands`, `edit`, `project`},
		`status`:    {`tmux`, `update`},
		`offscreen`: {`chat`, `!offscreen`},
		`info`:      {`twitch`, `bot`, `commands`, `file`, `edit`},
		`sync`:      {`twitch`, `bot`, `commands`, `sync`},
		`work`:      {`go`, `work`},
		`chat`:      {`twitch`, `chat`},
		`afk`:       {`twitch`, `chat`, `!afk`},
		`isosec`:    {`uniq`, `isosec`},
		`isonan`:    {`uniq`, `isonan`},
		`isodate`:   {`uniq`, `isodate`},
		`uuid`:      {`uniq`, `uuid`},
		`epoch`:     {`uniq`, `second`},
		`path`:      {`env`, `get`, `path`},
		//`long version of path`: {`env`, `get`, `path`},
	},

	Description: `
		Hi, I'm rwxrob.tv and this {{cmd .Name }} is my Bonzai™ tree. I am
		slowly replacing all my shell scripts and other Go utilities with
		Bonzai branches that I graft into this {{cmd .Name}} command. You
		are welcome to play around with it, but please know that I am
		radically changing things *daily*.
		
		Also check out https://github.com/rwxrob/foo for a sample template
		Bonzai tree to get started on your own.
		
		Here's some random output from the Dynamic *ls* function piped to
		the builtin *indent* function using {{ "{{ ls | indent 4 }}" }} Go
		template syntax:
		
		{{ ls | indent 4 }}
		
		That was a verbatim block because of the indent.
		
		`,
}
