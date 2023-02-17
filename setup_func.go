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
)

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
