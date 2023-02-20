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
		Usage:  `start application`,
		Action: server_func,
	},

	{
		Name: "setup",
		Usage: `setup tasks
		... user     # create a new database user entry
		... ex       # create a new database with exercises
		... user-ex  # create a new database user entries for exercises
		`,
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

			{
				Name:   "user-ex",
				Usage:  "create a new database user entries for exercises",
				Action: setup_user_ex_func,
			},
		},
	},
}

var app = &cli.App{
	Name: "one-o-one",
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
