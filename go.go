package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/fs/dir"
	"github.com/rwxrob/fs/file"
	"github.com/rwxrob/help"
	"github.com/rwxrob/term"
	"gopkg.in/yaml.v3"
)

var golang = &Z.Cmd{
	Name:     `go`,
	Summary:  `go related helper actions`,
	MinArgs:  1,
	Commands: []*Z.Cmd{help.Cmd, gowork, genisrune, sh2slice, godist},
}

var gowork = &Z.Cmd{
	Name:     `work`,
	Summary:  `turn on or off go.work file`,
	Usage:    `(on|off)`,
	MinArgs:  1,
	Params:   []string{"on", "off"},
	Commands: []*Z.Cmd{help.Cmd},
	Call: func(x *Z.Cmd, args ...string) error {
		switch args[0] {
		case "on":
			if file.Exists("go.work.off") {
				log.Print("go.work.off -> go.work")
				return os.Rename("go.work.off", "go.work")
			}
		case "off":
			if file.Exists("go.work") {
				log.Print("go.work -> go.work.off")
				return os.Rename("go.work", "go.work.off")
			}
		default:
			return x.UsageError()
		}
		return nil
	},
}

var genisrune = &Z.Cmd{
	Name:     `genisrune`,
	Summary:  `generate a performant rune check condition`,
	Commands: []*Z.Cmd{help.Cmd},
	MinArgs:  1,
	Call: func(x *Z.Cmd, args ...string) error {
		var conds []string
		for _, r := range args[0] {
			conds = append(conds, fmt.Sprintf("r==%q", r))
		}
		fmt.Println(strings.Join(conds, "||"))
		return nil
	},
}

var sh2slice = &Z.Cmd{
	Name:     `sh2slice`,
	Summary:  `splits a shell command into arguments`,
	Commands: []*Z.Cmd{help.Cmd},
	Call: func(_ *Z.Cmd, args ...string) error {
		list := []string{}
		if len(args) == 0 {
			args = append(args, term.Read())
		}
		for _, a := range args {
			// FIXME add awareness or globs and quoted segments
			for _, aa := range strings.Fields(a) {
				list = append(list, fmt.Sprintf("%q", aa))
			}
		}
		fmt.Println(strings.Join(list, ","))
		return nil
	},
}

func _build(args ...string) error {
	a := []string{`go`, `build`}
	a = append(a, args...)
	return Z.Exec(a...)
}

var godist = &Z.Cmd{
	Name:     `dist`,
	Summary:  `go distribution related commands`,
	Commands: []*Z.Cmd{godistbuild},
}

type GoBuildParams struct {
	Targets []GoBuildTarget
	O       map[string]any `yaml:",inline"`
}

type GoBuildTarget struct {
	OS   string
	Arch []string
}

var godistbuild = &Z.Cmd{
	Name:     `build`,
	Summary:  `build for for multiple architectures into dist dir`,
	Commands: []*Z.Cmd{help.Cmd},
	Description: `
			This build looks for a build.yaml file in the current directory
			and runs the build command on each building them all concurrently
			into the *dist* directory where they are ready for upload to
			GitHub as a release. Just add a README.md and run release.
	`,
	Call: func(_ *Z.Cmd, args ...string) error {
		if !file.Exists(`build.yaml`) {
			return _build(args...)
		}
		buf, err := os.ReadFile(`build.yaml`)
		if err != nil {
			return err
		}
		p := new(GoBuildParams)
		if err := yaml.Unmarshal(buf, p); err != nil {
			return err
		}
		os.RemoveAll(`dist`)
		dir.Create(`dist`)
		for _, target := range p.Targets {
			for _, arch := range target.Arch {
				log.Printf("Building for %v/%v", target.OS, arch)
				name := fmt.Sprintf("%v_%v_%v", dir.Name(), target.OS, arch)
				os.Setenv(`GOOS`, target.OS)
				os.Setenv(`GOARCH`, arch)
				_build(`-o`, `dist/`+name)
			}
		}
		return nil
	},
}
