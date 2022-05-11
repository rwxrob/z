package git

import (
	"fmt"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
	"github.com/rwxrob/term"
	"github.com/rwxrob/to"
)

func tags() ([]string, error) {
	lines := to.Lines(Z.Out("git", "tag"))
	return lines, nil
}

func deltag(tag string) error {
	if err := Z.Exec(`git`, `push`, `--delete`, `origin`, tag); err != nil {
		return err
	}
	return Z.Exec(`git`, `tag`, `--delete`, tag)
}

var Cmd = &Z.Cmd{
	Name:     `git`,
	Summary:  `git extensions`,
	Commands: []*Z.Cmd{help.Cmd, delTagsCmd},
}

var delTagsCmd = &Z.Cmd{
	Name:     `deltags`,
	Summary:  `delete all local and remote tags`,
	Commands: []*Z.Cmd{help.Cmd},
	Call: func(_ *Z.Cmd, args ...string) error {
		lines, err := tags()
		if err != nil {
			return err
		}
		fmt.Println(lines)
		if term.Prompt("Really delete them all? (y/N) ") != "y" {
			return nil
		}
		for _, tag := range lines {
			deltag(tag)
		}
		return err
	},
}
