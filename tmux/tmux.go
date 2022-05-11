package tmux

import (
	"os"
	"path"
	"strings"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
	"github.com/rwxrob/vars"
)

var Cmd = &Z.Cmd{
	Name:     `tmux`,
	Summary:  `make tmux updates`,
	Commands: []*Z.Cmd{help.Cmd, vars.Cmd, updateCmd, inCmd},
}

var inCmd = &Z.Cmd{
	Name:     `in`,
	Summary:  `exec a nested tmux session (unset TMUX)`,
	Usage:    `[help|<tmuxarg>...]`,
	Commands: []*Z.Cmd{help.Cmd},
	Call: func(_ *Z.Cmd, args ...string) error {
		home, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		conf := path.Join(home, `.tmux.conf`)
		tmuxargs := []string{`tmux`, `-f`, conf, `-u`}
		tmuxargs = append(tmuxargs, args...)
		os.Unsetenv(`TMUX`)
		Z.SysExec(tmuxargs...)
		return nil
	},
}

var updateCmd = &Z.Cmd{
	Name:    `update`,
	Summary: `update the onscreen status`,
	Call: func(_ *Z.Cmd, args ...string) error {
		msg := strings.Join(args, " ")
		return Z.Exec(
			"tmux", "-L", "live", "set", "-g", "status-right", msg,
		)
	},
}
