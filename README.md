# Personal Bonzai Commandz

[![GoDoc](https://godoc.org/github.com/rwxrob/cmds?status.svg)](https://godoc.org/github.com/rwxrob/z)
[![License](https://img.shields.io/badge/license-Apache2-brightgreen.svg)](LICENSE)

These days I prefer to maintain a single Go monolith command utility
rather than a ton of shell scripts in whatever languages. In fact, I
created [Bonzai](https://github.com/rwxrob/bonzai) specifically for this
sort of thing. I just `curl` down a single binary to whatever system I'm
on and I have all of my favorite functionality on *any* device with zero
compatibility hassles and installation dependencies. Everything just
works, *anywhere*. To update just make sure you have an Internet
connection and `z update`.

## Install

Note: you'll have to remove `go.work` yourself if you want to use this
build yourself (or clone all my structure the same way). I'm saving my
`go.work` in this Go repo only because it is a personal build.

```
go install github.com/rwxrob/z@latest
```

I prefer to use `z` instead of setting up a multicall binary since the
habits it builds into my muscle memory work on any operating system or
device and it doesn't take too much space when using UNIX pipelines and
such:

```
echo $(z isosec) $(z yaml2json quotes.yaml | jq -r .mad )
```

## Tab Completion

To activate bash completion just use the `complete -C` option from your
`.bashrc` or command line. There is no messy sourcing required. All the
completion is done by the program itself.

```
complete -C z z
```

If you don't have bash or tab completion check use the shortcut
commands instead.

## Embedded Documentation

All documentation (like manual pages) has been embedded into the source
code of the application. See the source or run the program with help to
access it.
