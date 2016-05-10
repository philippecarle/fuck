package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

var GlobalFlags = []cli.Flag{
	cli.StringFlag{
		EnvVar: "ENV_FUCK_NAME",
		Name:   "me, m",
		Usage:  "who's saying fuck off (default: your github username)",
	},
	cli.StringFlag{
		Name:  "who, w",
		Usage: "who's fucked off",
	},
}

var Commands = []cli.Command{}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
