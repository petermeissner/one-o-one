package main

import (
	"fmt"
	"log"
	"os"

	gbc "github.com/petermeissner/golang-basic-cred/library"
	"github.com/urfave/cli/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	// . "github.com/petermeissner/one-o-one/library"
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

				Name:  "setup",
				Usage: "setup application",
				Subcommands: []*cli.Command{

					{
						Name:  "ex",
						Usage: "create a new database with exercises",
						Action: func(cCtx *cli.Context) error {

							fmt.Println("Setting up Exercises in DB ...")

							// establish db connection
							gorm_dialect := sqlite.Open("gorm.db")
							db, err := gorm.Open(gorm_dialect, &gorm.Config{})
							if err != nil {
								log.Panic(err)
							}

							// ensure tables exist
							db.AutoMigrate(&Exercise{})

							// add exercise data
							m := h_generate_multiply()
							db.Clauses(clause.OnConflict{
								Columns:   []clause.Column{{Name: "exercise"}},
								UpdateAll: true,
							}).Create(&m)

							d := h_generate_division()
							db.Clauses(clause.OnConflict{
								Columns:   []clause.Column{{Name: "exercise"}},
								UpdateAll: true,
							}).Create(&d)

							return err
						},
					},

					{
						Name:  "user",
						Usage: "create a new database user entry",
						Action: func(cCtx *cli.Context) error {
							fmt.Println("User creation ...")

							auth := gbc.Get_auth_from_term()

							// establish db connection
							gorm_dialect := sqlite.Open("gorm.db")
							db, err := gorm.Open(gorm_dialect, &gorm.Config{})
							if err != nil {
								log.Panic(err)
							}

							// Hash credentials and store them in db
							gbc.Upsert_auth_as_credential_to_db(db, auth)

							return err
						},
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
