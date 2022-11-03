package env

import (
	"fmt"
	"os"
	"strings"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
	"github.com/rwxrob/term"
)

var Cmd = &Z.Cmd{
	Name:     `env`,
	Summary:  `commands for environment variables`,
	Commands: []*Z.Cmd{getCmd, dataCmd, help.Cmd},
}

var dataCmd = &Z.Cmd{
	Name:     `data`,
	Aliases:  []string{`all`},
	Summary:  `print environment data to stdout`,
	Commands: []*Z.Cmd{help.Cmd},
	Call: func(_ *Z.Cmd, _ ...string) error {
		for _, pair := range os.Environ() {
			fmt.Println(pair)
		}
		return nil
	},
}

var getCmd = &Z.Cmd{
	Name:     `get`,
	Usage:    `(help|NAME)`,
	Summary:  `print specified environment variable to stdout`,
	Commands: []*Z.Cmd{help.Cmd},
	NumArgs:  1,
	Call: func(_ *Z.Cmd, args ...string) error {
		v := os.Getenv(args[0])
		if v == "" {
			v = os.Getenv(strings.ToUpper(args[0]))
		}
		_, err := term.Print(v)
		return err
	},
}
