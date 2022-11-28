# 🌳 Personal Bonzai Commandz

These days I prefer to maintain a single Go [stateful command tree monolith](https://rwxrob.github.io/zet/1729/) tool rather than a ton of shell scripts in whatever languages. In fact, I created [Bonzai](https://github.com/rwxrob/bonzai) specifically for this sort of thing. I just `curl` down a single binary to whatever system I'm on and I have all of my favorite functionality on *any* device with zero compatibility hassles and installation dependencies. Everything just works, *anywhere*.

## Install

Just download one of the [release binaries](https://github.com/rwxrob/z/releases):

```
curl -L https://github.com/rwxrob/z/releases/latest/download/z-linux-amd64 -o ~/.local/bin/rwxrobz
curl -L https://github.com/rwxrob/z/releases/latest/download/z-darwin-amd64 -o ~/.local/bin/rwxrobz
curl -L https://github.com/rwxrob/z/releases/latest/download/z-darwin-arm64 -o ~/.local/bin/rwxrobz
curl -L https://github.com/rwxrob/z/releases/latest/download/z-windows-amd64 -o ~/.local/bin/rwxrobz
```

Or install directly with `go`:

```
go install github.com/rwxrob/z@latest
```

Note: you'll have to remove `go.work` yourself if you want to use this build yourself (or clone all my structure the same way). I'm saving my `go.work` in this Go repo only because it is a personal build.

I prefer to use `z` instead of setting up a multicall binary since the habits it builds into my muscle memory work on any operating system or device and it doesn't take too much space when using UNIX pipelines and
such:

```
echo $(z isosec) $(z y2j quotes.yaml | jq -r .mad )
```

## Tab Completion

To activate bash completion just use the `complete -C` option from your `.bashrc` or command line. There is no messy sourcing required. All the completion is done by the program itself.

```
complete -C z z
```

If you don't have bash or tab completion check use the shortcut commands instead. Zsh does a good job of learning your commands over time all by itself, but some of the custom completions may not work as well. Personally, I prefer the default Linux shell (Bash) over the default Mac shell (Zsh). (PRs welcome to integrate completion into Zsh without dumping a ton of shell code that has to be sourced.)

## Embedded Documentation

All documentation (like manual pages) has been embedded into the source code of the application. See the source or run the program with help to access it.

## Building

Releases are built using the following commands:

```
z go build
gh release create
gh release upload TAG build/*
```
