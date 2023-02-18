package lib

import (
	"fmt"
)

type Exercise struct {
	Exercise string `gorm:"primaryKey"`
	Solution float32
}

type Exercised struct {
	exercise_id uint
	user_id     uint
	runs        uint
	fails       uint
	success     uint
	duration_s  uint
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

func H_generate_multiply() ExerciseSlice {
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

func H_generate_division() ExerciseSlice {
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
