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
	app.Get("/fraud", HandleFraud)
	app.Get("/security", HandleSecurity)
	app.Get("/b2b", HandleB2B)
	app.Get("/card", HandleCard)
	app.Get("/payment", HandlePayment)
	app.Get("/wallet", HandleWallet)
	app.Get("/payments", HandlePayment)
	app.Get("/exchange", HandleExchange)
	app.Get("/faq", HandleFAQ)
	/* Page Not und Management */

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendFile("./static/page-error-404.html")
	})
}
