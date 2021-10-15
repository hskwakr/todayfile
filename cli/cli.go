package cli

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/hskwakr/todayfile/todayfile"
)

const (
	ApplicationName = "todayfile"

	ExitCodeOK               = 0
	ExitCodeParseFlagError   = 1
	ExitCodeArgumentsError   = 1
	ExitCodeApplicationError = 1
)

type CLI struct {
	OutStream, ErrStream io.Writer
}

func (c *CLI) Run(args []string) int {
	if r := c.parse(args); r != ExitCodeOK {
		return r
	}

	err := todayfile.Create()
	if err != nil {
		log.Println(err)

		return ExitCodeApplicationError
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
