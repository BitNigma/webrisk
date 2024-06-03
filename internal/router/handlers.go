package router

import (
	"websirk/internal/views/index"

	"github.com/gofiber/fiber/v2"
)

type Auth struct {
	Islogin    bool
	Email      string
	KYC        bool
	ErrorPass  error
	ErrorEmail error
}

// Handle main page
func HandleHome(c *fiber.Ctx) error {
	c.Set("Content-type", "text/html")
	return index.MainPage().Render(c.UserContext(), c.Context().Response.BodyWriter())
}

func HandleAbout(c *fiber.Ctx) error {
	c.Set("Content-type", "text/html")
	return index.AboutPage().Render(c.UserContext(), c.Context().Response.BodyWriter())
}

func HandleContacts(c *fiber.Ctx) error {
	c.Set("Content-type", "text/html")
	return index.ContactsPage().Render(c.UserContext(), c.Context().Response.BodyWriter())
}

func HandleTerms(c *fiber.Ctx) error {
	c.Set("Content-type", "text/html")
	return index.TermsPage().Render(c.UserContext(), c.Context().Response.BodyWriter())
}

func HandleCard(c *fiber.Ctx) error {
	return c.Render("./static/terms.html", fiber.Map{
		"Jonny": "Hello, World!",
	})
}

func HandleWallet(c *fiber.Ctx) error {
	c.Set("Content-type", "text/html")
	return index.WalletPage().Render(c.UserContext(), c.Response().BodyWriter())
}

func HandlePayment(c *fiber.Ctx) error {
	return c.Render("./static/terms.html", fiber.Map{
		"Jonny": "Hello, World!",
	})
}

func HandlePolicy(c *fiber.Ctx) error {
	c.Set("Content-type", "text/html")
	return index.PrivacyHanle().Render(c.UserContext(), c.Response().BodyWriter())
}

func HandleAML(c *fiber.Ctx) error {
	c.Set("Content-type", "text/html")
	return index.AMLHanle().Render(c.UserContext(), c.Response().BodyWriter())
}

func HandleB2C(c *fiber.Ctx) error {
	return c.Render("./static/terms.html", fiber.Map{
		"Jonny": "Hello, World!",
	})
}

func HandleB2B(c *fiber.Ctx) error {
	return c.Render("./static/terms.html", fiber.Map{
		"Jonny": "Hello, World!",
	})
}
