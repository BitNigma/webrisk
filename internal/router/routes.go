package router

import (
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	// Static files
	app.Static("/static", "./static")

	/* Views */
	app.Get("/", HandleHome)
	app.Get("/about", HandleAbout)
	app.Get("/contacts", HandleContacts)
	app.Get("/terms", HandleTerms)
	app.Get("/privacy", HandlePolicy)
	app.Get("/fraud", HandleAML)
	app.Get("/b2b", HandleB2C)
	app.Get("/b2b", HandleB2B)
	app.Get("/card", HandleCard)
	app.Get("/payment", HandlePayment)
	app.Get("/wallet", HandleWallet)
	/* Page Not und Management */

	/* Page Not Found Management */

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendFile("./static/page-error-404.html")
	})
}
