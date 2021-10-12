package cli

import (
	"flag"
	"fmt"
	"io"
	"os"
)

const (
	ApplicationName = "todayfile"

	ExitCodeOK               = 0
	ExitCodeParseFlagError   = 0
	ExitCodeArgumentsError   = 0
	ExitCodeApplicationError = 0
)

type CLI struct {
	OutStream, ErrStream io.Writer
	args
}

type args struct {
}

func (c *CLI) Run(args []string) int {
	if r := c.parse(args); r != ExitCodeOK {
		return r
	}

	return ExitCodeOK
}

func (c *CLI) parse(args []string) int {
	flags := flag.NewFlagSet(ApplicationName, flag.ContinueOnError)
	flags.SetOutput(c.ErrStream)

	flags.Usage = func() {
		fmt.Fprintf(
			flag.CommandLine.Output(),
			"Create afile with the name of today's date.\n\nUsage:\n\t"+ApplicationName+" \n\n",
		)
		flags.PrintDefaults()
		os.Exit(0)
	}

	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeParseFlagError
	}

	return ExitCodeOK
}
