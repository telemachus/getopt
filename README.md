# getopt (forked from rsc.io/getopt)

## Note

This repository is a fork of [rsc's][rsc] [getopt][rgetopt].  I have
updated the code in various ways, and I will probably continue to do so.  But
the core idea (and implementation) is his rather than mine. td;dr — credit where
credit is due, but all mistakes and bad ideas are my fault and not rsc's.

[rsc]: https://github.com/rsc
[rgetopt]: https://github.com/rsc/getopt

## Introduction

Package getopt parses command lines using [`getopt_(3)`][getopt3] syntax.  It is
a replacement for [`FlagSet.Parse`][fs.Parse] but still expects flags themselves
to be defined in package `flag`.

[getopt3]: http://man7.org/linux/man-pages/man3/getopt.3.html
[fs.Parse]: https://pkg.go.dev/flag#FlagSet.Parse

Flags defined with one-letter names are available as short flags (invoked using
one dash, as in `-x`) and all flags are available as long flags (invoked using
two dashes, as in `--x` or `--xylophone`).

To use, define a [`FlagSet`][FlagSet] and flags as usual with [package
flag][flag].  Then introduce any aliases by calling `getopt.Alias`:

```go
type cfg struct {
	quiet bool
	verbose  bool
}

func main() {
	fs := getopt.NewFlagSet("awesome", flag.ContinueOnError)
	cfg := &cfg{}

	fs.BoolVar(&cfg.quiet, "quiet", false, "Print only errors")
	fs.BoolVar(&cfg.help, "v", false, "Print way too much")

	getopt.Alias("v", "verbose")
	getopt.Alias("q", "quiet")
}
```

Or call `getopt.Aliases` to define a list of aliases:

```go
getopt.Aliases(
	"q", "quiet",
	"v", "verbose",
)
```

[FlagSet]: https://pkg.go.dev/flag#FlagSet
[flag]: https://godoc.org/flag

One name in each pair must already be defined in package `flag` (so either
"q" or "quiet", and also either "v" or "verbose").

Then parse the command-line using [`fs.Parse()`][fs.Parse].

When writing a custom `flag.Usage` function, call `getopt.PrintDefaults` instead
of `flag.PrintDefaults` to get a usage message that includes the names of
aliases in flag descriptions.
