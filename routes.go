package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

// route: root
//
// default html page rendering content
func route_root(c *fiber.Ctx) error {

	sess, err := store.Get(c)
	if err != nil {
		log.Println(err)
	}

	// log.Println(sess)
	name := sess.Get("name")
	isLogin := sess.Get("is_logged_in")
	UnauthorizedMessage := "Kein Login, keine Kekse."
	AuthorizedMessage := fmt.Sprintf("Willkommen %v", name)

	return c.Render("index", fiber.Map{
		"AuthorizedMessage":   AuthorizedMessage,
		"UnauthorizedMessage": UnauthorizedMessage,
		"IsLogin":             isLogin,
	})

}

// route: login
//
// executes and redirects to root
func route_login(c *fiber.Ctx) error {

	// get/create session
	sess, err := store.Get(c)
	if err != nil {
		panic(err)
	}
	logger_info.Printf("SessionID: %+v\n", sess.ID())
	logger_info.Printf("SessionKeys%+v\n", sess.Keys())

	// parse authentication information
	auth := new(Auth)
	if err := c.BodyParser(auth); err != nil {
		return err
	}

	// check credentials and set session values
	login_handler(auth, sess)

	return c.Redirect("/")

}

// route: logout
//
// executes and redirects to root
func route_logout(c *fiber.Ctx) error {
	// get/create session
	sess, err := store.Get(c)
	if err != nil {
		panic(err)
	}

	// handle logout
	sess.Delete("name")
	if err := sess.Destroy(); err != nil {
		panic(err)
	}

	// return
	return c.Redirect("/")
}

// route: answer
//
// executes and redirects to root
func route_answer(c *fiber.Ctx) error {

	// get/create session
	sess, err := store.Get(c)
	if err != nil {
		panic(err)
	}
	logger_info.Printf("SessionID: %+v\n", sess.ID())
	logger_info.Printf("SessionKeys%+v\n", sess.Keys())

	// check credentials and set session values
	answer_handler(sess)

	return c.Redirect("/")

}

// Setup setup_routes
//
// function adding routes to server app
// route: root, for static content
func setup_routes(app *fiber.App) {
	app.Static("/", "./public")

	// route: root
	app.Get("/", route_root)

	// route: login
	// executes and redirects to root
	app.Post("/login", route_login)

	// route: logout
	// executes and redirects to root
	app.Get("/logout", route_logout)

	// route: answer
	// executes and redirects to root
	app.Post("/answer", route_answer)
}
