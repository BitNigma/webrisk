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

func HandlePolicy(c *fiber.Ctx) error {
	c.Set("Content-type", "text/html")
	return index.PrivacyHanle().Render(c.UserContext(), c.Response().BodyWriter())
}

func HandleWallet(c *fiber.Ctx) error {
	c.Set("Content-type", "text/html")
	return index.WalletPage().Render(c.UserContext(), c.Response().BodyWriter())
}

func HandleSecurity(c *fiber.Ctx) error {
	c.Set("Content-type", "text/html")
	return index.SecurityHanle().Render(c.UserContext(), c.Response().BodyWriter())
}

func HandleFraud(c *fiber.Ctx) error {
	c.Set("Content-type", "text/html")
	return index.FraudHanle().Render(c.UserContext(), c.Response().BodyWriter())
}

func HandleAPI(c *fiber.Ctx) error {
	c.Set("Content-type", "text/html")
	return index.APIHandle().Render(c.UserContext(), c.Response().BodyWriter())
}

func HandleCompliance(c *fiber.Ctx) error {
	c.Set("Content-type", "text/html")
	return index.HandleCompliance().Render(c.UserContext(), c.Response().BodyWriter())
}

func HandlePayment(c *fiber.Ctx) error {
	c.Set("Content-type", "text/html")
	return index.PaymentPage().Render(c.UserContext(), c.Response().BodyWriter())
}

func HandleB2B(c *fiber.Ctx) error {
	c.Set("Content-type", "text/html")
	return index.B2BPage().Render(c.UserContext(), c.Response().BodyWriter())
}

func HandleExchange(c *fiber.Ctx) error {
	c.Set("Content-type", "text/html")
	return index.ExchangePage().Render(c.UserContext(), c.Response().BodyWriter())
}

func HandleMessage(c *fiber.Ctx) error {
	c.Set("Content-type", "text/html")
	return index.WalletPage().Render(c.UserContext(), c.Response().BodyWriter())
}

func HandleFAQ(c *fiber.Ctx) error {
	return c.Render("./static/page-faq.html", fiber.Map{
		"Jonny": "Hello, World!",
	})
}
