package main

import (
	"github.com/gofiber/fiber/v2/middleware/session"
)

// # login_handler
// - called for side effects on session
//
// - check if credentials are ok
//
// - set appropriate values in session variables
func answer_handler(sess *session.Session) {

	// // establish db connection
	// gorm_dialect := sqlite.Open("gorm.db")
	// db, err := gorm.Open(gorm_dialect, &gorm.Config{})
	// if err != nil {
	// 	log.Panic(err)
	// }

	// check credentials
	logger_info.Println(sess.Get("is_logged_in").(string))

	// // handle session according to authentication status
	// if user_known {

	// 	sess.Set("name", auth.Name)
	// } else {
	// 	sess.Set("is_logged_in", false)
	// 	sess.Set("name", nil)
	// }

	if err := sess.Save(); err != nil {
		panic(err)
	}
}
