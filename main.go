package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/philippecarle/fuck/actions"
)

func main() {

	app := cli.NewApp()
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
