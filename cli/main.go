package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Exercise struct {
	Exercise string `gorm:"primaryKey"`
	Solution float32
}

type Exercised struct {
	excessive_id uint
	player_id    uint
	runs         uint
	fails        uint
	success      uint
	duration_s   uint
}

func (e Exercise) print_all() {
	fmt.Println(e.Exercise, "=", e.Solution)
}

func (e Exercise) print_ex() {
	fmt.Println(e.Exercise, "= ?")
}

type ExerciseSlice []Exercise

func (e ExerciseSlice) print_it() {
	for _, s := range e {
		s.print_all()
	}
}

func h_generate_multiply() ExerciseSlice {
	var e = make(ExerciseSlice, 100, 100)
	c := 0
	for a := 1; a <= 10; a++ {
		for b := 1; b <= 10; b++ {
			e[c] = Exercise{Exercise: "" + fmt.Sprint(a) + " * " + fmt.Sprint(b), Solution: float32(a * b)}
			c = c + 1
		}
	}
	return e
}

func h_generate_division() ExerciseSlice {
	var e = make(ExerciseSlice, 100, 100)
	c := 0
	for a := 1; a <= 10; a++ {
		for b := 1; b <= 10; b++ {
			e[c] = Exercise{Exercise: "" + fmt.Sprint(a*b) + " / " + fmt.Sprint(a), Solution: float32(b)}
			c = c + 1
		}
	}
	return e
}

func h_generate_subtraction() ExerciseSlice {
	var e = make(ExerciseSlice, 100, 100)
	return e
}

func h_generate_addition() ExerciseSlice {
	var e = make(ExerciseSlice, 100, 100)
	return e
}

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
							d := h_generate_division()
							d.print_it()

							m := h_generate_multiply()
							m.print_it()

							// establish db connection
							gorm_dialect := sqlite.Open("gorm.db")
							db, err := gorm.Open(gorm_dialect, &gorm.Config{})
							if err != nil {
								log.Panic(err)
							}

							db.AutoMigrate(&Exercise{})
							db.Clauses(clause.OnConflict{
								Columns:   []clause.Column{{Name: "exercise"}}, // key colume
								UpdateAll: true,
							}).Create(&m)
							// db.Create(&m)

							return err
						},
					},
					{
						Name:  "user",
						Usage: "create a new database user entry",
						Action: func(cCtx *cli.Context) error {
							fmt.Println("user: ", cCtx.Args().First())
							return nil
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
