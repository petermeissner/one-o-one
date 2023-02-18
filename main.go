package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

// The default logger to be (re-)used throughout the application
var logger_info = log.Default()

var commands = []*cli.Command{
	{
		Name:   "server",
		Usage:  "run server",
		Action: server_func,
	},

	{
		Name:  "setup",
		Usage: "setup application",
		Subcommands: []*cli.Command{
			{
				Name:   "ex",
				Usage:  "create a new database with exercises",
				Action: setup_ex_func,
			},

			{
				Name:   "user",
				Usage:  "create a new database user entry",
				Action: setup_user_func,
			},
		},
	},
}

var app = &cli.App{
	Name:  "one-o-one",
	Usage: "set up ono-o-one server",
	Action: func(*cli.Context) error {
		fmt.Println("Welcome to one-o-one-cli.")
		fmt.Println("use --help to get help")

		return nil
	},
	Commands: commands,
}

func main() {

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
