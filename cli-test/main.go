package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "cli-test",
		Usage: "set up ono-o-one server",
		Action: func(*cli.Context) error {
			fmt.Println("Welcome to ooo-cli.")
			fmt.Println("use --help to get help")
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:  "db-ex",
				Usage: "create a new database with exercises",
				Action: func(cCtx *cli.Context) error {
					return nil
				},
			},
			{
				Name:  "db-user",
				Usage: "create a new database user entry",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("user: ", cCtx.Args().First())
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
