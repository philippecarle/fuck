package main

import (
	"github.com/philippecarle/fuck/actions"
	"os"
	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()

	// app informations
	app.Name = Name
	app.Version = Tag
	app.Author = "philippecarle"
	app.Email = "tr@nsfer.red"
	app.Usage = "Wanna tell fuck off to anybody without leaving the CLI ?"

	app.Flags = GlobalFlags

	app.Commands = Commands
	app.CommandNotFound = CommandNotFound

	app.Action = actions.Fuck

	app.Setup()

	app.Run(os.Args)
}
