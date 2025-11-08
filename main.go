package main

import (
	"os"

	"github.com/alecthomas/kong"
	"github.com/jmilne22/distrohop/cli"
)

type CLI struct {
	cli.Globals
	State cli.StateCmd `cmd:"" help:"Check state of installed packages"`
}

func main() {
	cli := CLI{}

	if len(os.Args) < 2 {
		os.Args = append(os.Args, "--help")
	}

	ctx := kong.Parse(&cli,
		kong.Name("distrohop"),
		kong.Description("A somewhat distro agnostic tool to bootstrap new systems for terminal distrohoppers"),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
		}),
	)
	err := ctx.Run(&cli.Globals)
	ctx.FatalIfErrorf(err)
}
