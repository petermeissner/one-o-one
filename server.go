package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/urfave/cli/v2"
)

var store *session.Store

type Message struct {
	Message string
}

// Server definition
//
// Defines server config and endpoints
func server_func(cCtx *cli.Context) error {

	// new session-store
	store = session.New()

	app := fiber.New()

	// route: root
	app.Get("/", func(c *fiber.Ctx) error {

		sess, err := store.Get(c)
		if err != nil {
			log.Println(err)
		}

		// log.Println(sess)
		isLogin := sess.Get("is_logged_in")

		if isLogin == true {
			return c.JSON(Message{Message: "logged in"})
		} else {
			return c.JSON(Message{Message: "NOT logged in"})
		}
	})

	// route: login
	// executes and redirects to root
	app.Post("/login", func(c *fiber.Ctx) error {
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
	})

	// route: logout
	// executes and redirects to root
	app.Get("/logout", func(c *fiber.Ctx) error {
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
	})

	// start and listen
	log.Fatal(app.Listen("localhost:3000"))

	// adhere to cli function signature requirements
	return nil
}
