package main

import (
	"fmt"
	"github.com/aozora0000/paste-master/command"
	"github.com/codegangsta/cli"
	"os"
)

var (
	name    = "paste-master"
	version = "dev"
	commit  = "none"
)

func main() {
	app := cli.NewApp()
	app.Name = name
	app.Version = fmt.Sprintf("%s-%s", version, commit)
	app.Author = "aozora0000"
	app.Email = "aozora0000@gmail.com"
	app.Usage = "Filtering Paste From ClipBoard."

	app.Flags = command.Flags
	app.Action = command.Command
	app.CommandNotFound = CommandNotFound

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
