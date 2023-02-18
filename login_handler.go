package main

import (
	"log"

	"github.com/gofiber/fiber/v2/middleware/session"
	gbc "github.com/petermeissner/golang-basic-cred/library"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Auth parse type
type Auth struct {
	Name     string `json:"name" xml:"name" form:"name"`
	Password string `json:"password" xml:"password" form:"password"`
}

// # login_handler
// - called for side effects on session
//
// - check if credentials are ok
//
// - set appropriate values in session variables
func login_handler(auth *Auth, sess *session.Session) {

	// establish db connection
	gorm_dialect := sqlite.Open("gorm.db")
	db, err := gorm.Open(gorm_dialect, &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}

	// check credentials
	gbc_auth := gbc.Auth{Username: auth.Name, Password: auth.Password}
	user_known := gbc.Check_credential(db, gbc_auth)
	logger_info.Printf("user: %+v, login success: %+v\n", auth.Name, user_known)

	// handle session according to authentication status
	if user_known {
		sess.Set("is_logged_in", true)
		sess.Set("name", auth.Name)
	} else {
		sess.Set("is_logged_in", false)
		sess.Set("name", nil)
	}

	if err := sess.Save(); err != nil {
		panic(err)
	}
}
