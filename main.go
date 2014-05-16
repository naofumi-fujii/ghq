package main

import (
	"os"

	"github.com/codegangsta/cli"
)

var VERSION string

func main() {
	app := cli.NewApp()
	app.Name = "ghq"
	app.Usage = "Manage GitHub repository clones"
	app.Version = VERSION
	app.Author = "motemen"
	app.Email = "motemen@gmail.com"
	app.Commands = Commands
	app.Run(os.Args)
}