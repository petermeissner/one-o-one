package main

import (
	"fmt"
	gbc "github.com/petermeissner/golang-basic-cred/library"
	oool "github.com/petermeissner/one-o-one/lib"
	"github.com/urfave/cli/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"os"
)

func server_func(cCtx *cli.Context) error {
	fmt.Println("SERVER!!!")
	return nil
}

func setup_ex_func(cCtx *cli.Context) error {

	fmt.Println("Setting up Exercises in DB ...")

	// establish db connection
	gorm_dialect := sqlite.Open("gorm.db")
	db, err := gorm.Open(gorm_dialect, &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}

	// ensure tables exist
	db.AutoMigrate(&oool.Exercise{})

	// add exercise data
	m := oool.H_generate_multiply()
	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "exercise"}},
		UpdateAll: true,
	}).Create(&m)

	d := oool.H_generate_division()
	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "exercise"}},
		UpdateAll: true,
	}).Create(&d)

	return err
}

func setup_user_func(cCtx *cli.Context) error {
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
}

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

func main() {
	app := &cli.App{
		Name:  "one-o-one",
		Usage: "set up ono-o-one server",
		Action: func(*cli.Context) error {
			fmt.Println("Welcome to ooo-cli.")
			fmt.Println("use --help to get help")

			return nil
		},
		Commands: commands,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
