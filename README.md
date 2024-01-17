# 🌳 Personal Commandz

These days I prefer to maintain a single Go [stateful command tree monolith](https://rwxrob.github.io/zet/1729/) tool rather than a ton of shell scripts in whatever languages. I just `curl` down a single binary to whatever system I'm on and I have all of my favorite functionality on *any* device with zero compatibility hassles and installation dependencies. Everything just works, *anywhere*.

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

I prefer to use `z` instead of setting up a multi-call binary since the habits it builds into my muscle memory work on any operating system or device and it doesn't take too much space when using UNIX pipelines and
such:

```
echo $(z isosec) $(z y2j quotes.yaml | jq -r .mad )
```

## Tab Completion

To activate bash completion just source the completion code in your favorite shell:

```bash
sourc <(z completion bash)
```
