# Personal Bonzai Command Tree (Monolith)

[![GoDoc](https://godoc.org/github.com/rwxrob/cmds?status.svg)](https://godoc.org/github.com/rwxrob/cmds)
[![License](https://img.shields.io/badge/license-Apache2-brightgreen.svg)](LICENSE)

These days I prefer to maintain a single Go monolith command utility
with everything I would have used shell scripts for before. I created
[Bonzai](https://github.com/rwxrob/bonzai) specifically for this sort of
thing. This way I just have to copy a single binary over to whatever
system I'm working on and I have all of my favorite functionality on
*any* device that Go supports with zero compatibility hassles and
installation dependencies. It just works, and by works I mean
*anywhere*. Hell, I don't even need a container (but can easily make a
FROM SCRATCH container with nothing but `cmds` in it). If I want a
subset of the commands I just trim the tree and compose them into a
different monolith --- in minutes.

## Install

```
go install github.com/rwxrob/cmds@latest
```

I have `z` hard link as well to keep things easy to type.

```
ln "$GOBIN/cmds" "$GOBIN/z"
```

## Tab Completion

To activate bash completion just use the `complete -C` option from your
`.bashrc` or command line. There is no messy sourcing required. All the
completion is done by the program itself.

```
complete -C cmds cmds
complete -C z z
```

If you don't have bash or tab completion check use the shortcut
commands instead.

## Embedded Documentation

All documentation (like manual pages) has been embedded into the source
code of the application. See the source or run the program with help to
access it.
