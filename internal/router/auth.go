package router

import (
	"github.com/gofiber/fiber/v2"
)

// Authentication Middleware
func AuthMiddleware(c *fiber.Ctx) error {
	cookie := c.Cookies("session_id")
	if len(cookie) < 1 {
		return c.Redirect("/login")
	}
	return c.Next()
}

func HandleLogout(c *fiber.Ctx) error {
	c.ClearCookie("session_id")
	return c.Redirect("/")
}
