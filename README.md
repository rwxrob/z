# Bonzai™ Sample `foo` Command (Template)

*Create a new GitHub project using this template and change this
README.md to match your project. Make all your template changes before
making your first commit.*

![WIP](https://img.shields.io/badge/status-wip-red)
![Go Version](https://img.shields.io/github/go-mod/go-version/rwxrob/foo)
[![GoDoc](https://godoc.org/github.com/rwxrob/foo?status.svg)](https://godoc.org/github.com/rwxrob/foo)
[![License](https://img.shields.io/badge/license-Apache2-brightgreen.svg)](LICENSE)

## Install

This command can be installed as a standalone program or composed into 
a Bonzai command tree.

Standalone

```
go install github.com/rwxrob/foo/foo@latest
```

Composed

```go
package cmds

import (
	"github.com/rwxrob/bonzai"
	"github.com/rwxrob/foo"
)

var Cmd = &bonzai.Cmd{
	Name:     `cmds`,
	Commands: []*bonzai.Cmd{help.Cmd, foo.Cmd},
}
```

## Tab Completion

To activate bash completion just use the `complete -C` option from your
`.bashrc` or command line. There is no messy sourcing required. All the
completion is done by the program itself.

```
complete -C foo foo
```

If you don't have bash or tab completion check use the shortcut
commands instead.

## Embedded Documentation

All documentation (like manual pages) has been embedded into the source
code of the application. See the source or run the program with help to
access it.
