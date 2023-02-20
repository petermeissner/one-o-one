package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/html"
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

	// initialize template engine
	engine := html.New("./views", ".html")
	engine.Reload(true)
	engine.Delims("{{", "}}")

	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
	})

	// register routes
	setup_routes(app)

	// start and listen
	log.Fatal(app.Listen("localhost:3000"))

	// adhere to cli function signature requirements
	return nil
}
