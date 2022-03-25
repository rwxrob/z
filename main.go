package main

import (
	Z "github.com/rwxrob/bonzai"
	"github.com/rwxrob/bonzai/inc/help"
	"github.com/rwxrob/config"
	"github.com/rwxrob/twitch"
	"github.com/rwxrob/yaml2json"
)

func main() {
	Cmd.Run()
}

var Cmd = &Z.Cmd{
	Name:      `z`,
	Summary:   `rwxrob's bonzai command tree`,
	Version:   `v0.0.1`,
	Copyright: `Copyright 2021 Robert S Muhlestein`,
	License:   `Apache-2.0`,
	Commands: []*Z.Cmd{
		help.Cmd, yaml2json.Cmd, twitch.Cmd, config.Cmd, tmux,
	},
}
