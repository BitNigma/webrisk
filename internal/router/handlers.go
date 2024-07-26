package router

import (
	"log"
	"time"
	"websirk/internal/model"
	"websirk/internal/userstore"
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
	return index.MainPage(c).Render(c.UserContext(), c.Context().Response.BodyWriter())
}

func HandleAbout(c *fiber.Ctx) error {
	c.Set("Content-type", "text/html")
	return index.AboutPage(c).Render(c.UserContext(), c.Context().Response.BodyWriter())
}

func HandleContacts(c *fiber.Ctx) error {
	c.Set("Content-type", "text/html")
	return index.ContactsPage(c).Render(c.UserContext(), c.Context().Response.BodyWriter())
}

func HandleTerms(c *fiber.Ctx) error {
	c.Set("Content-type", "text/html")
	return index.TermsPage(c).Render(c.UserContext(), c.Context().Response.BodyWriter())
}

func HandlePolicy(c *fiber.Ctx) error {
	c.Set("Content-type", "text/html")
	return index.PrivacyHanle(c).Render(c.UserContext(), c.Response().BodyWriter())
}

func HandleWallet(c *fiber.Ctx) error {
	c.Set("Content-type", "text/html")
	return index.WalletPage(c).Render(c.UserContext(), c.Response().BodyWriter())
}

func HandleSecurity(c *fiber.Ctx) error {
	c.Set("Content-type", "text/html")
	return index.SecurityHanle(c).Render(c.UserContext(), c.Response().BodyWriter())
}

func HandleFraud(c *fiber.Ctx) error {
	c.Set("Content-type", "text/html")
	return index.FraudHanle(c).Render(c.UserContext(), c.Response().BodyWriter())
}

func HandleAPI(c *fiber.Ctx) error {
	c.Set("Content-type", "text/html")
	return index.APIHandle(c).Render(c.UserContext(), c.Response().BodyWriter())
}

func HandleCompliance(c *fiber.Ctx) error {
	c.Set("Content-type", "text/html")
	return index.HandleCompliance(c).Render(c.UserContext(), c.Response().BodyWriter())
}

func HandlePayment(c *fiber.Ctx) error {
	c.Set("Content-type", "text/html")
	return index.PaymentPage(c).Render(c.UserContext(), c.Response().BodyWriter())
}

func HandleB2B(c *fiber.Ctx) error {
	c.Set("Content-type", "text/html")
	return index.B2BPage(c).Render(c.UserContext(), c.Response().BodyWriter())
}

func HandleExchange(c *fiber.Ctx) error {
	c.Set("Content-type", "text/html")
	return index.ExchangePage(c).Render(c.UserContext(), c.Response().BodyWriter())
}

func HandleMessage(c *fiber.Ctx) error {
	c.Set("Content-type", "text/html")
	return index.WalletPage(c).Render(c.UserContext(), c.Response().BodyWriter())
}

func HandleFAQ(c *fiber.Ctx) error {
	return c.Render("./static/page-faq.html", fiber.Map{
		"Jonny": "Hello, World!",
	})
}

// Handle POST request
func HandleLoginAuth(c *fiber.Ctx) error {

	cookie := c.Cookies("session_id")
	if len(cookie) > 1 {
		return c.Redirect("/personal/dash")
	}

	auth := &Auth{}
	sess, err := store.Get(c)
	if err != nil {
		log.Println("can't get session store", err)
		return err
	}

	//Create User
	user := model.User{
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
	}
	store := userstore.NewStore()
	if err = store.CheckAuthUser(user); err != nil {
		switch err {
		case userstore.ErrorByEmail:
			auth.ErrorEmail = err
		case userstore.ErrorPassword:
			auth.ErrorPass = err
		}
		c.Set("Content-type", "text/html")
		return index.HandleLogin(err.Error(), c).Render(c.UserContext(), c.Response().BodyWriter())
	}
	sess.Set("session_id", "super-test")
	if err = sess.Save(); err != nil {
		log.Println("can't save user session in store", err)
		return err
	}
	return c.Redirect("/personal/dash")
}

// Render Sign In page
func HandleLogin(c *fiber.Ctx) error {

	cookie := c.Cookies("session_id")
	if len(cookie) > 1 {
		return c.Redirect("/personal/dash")
	}

	c.Set("Content-type", "text/html")
	return index.HandleLogin("", c).Render(c.UserContext(), c.Response().BodyWriter())

}

// Handle Post request to Create New User
func HandleRegister(c *fiber.Ctx) error {

	cookie := c.Cookies("session_id")
	if len(cookie) > 1 {
		return c.Redirect("/personal/dash")
	}

	sess, err := store.Get(c)
	if err != nil {
		log.Println(err)
	}

	//Create User
	user := model.User{
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
		Balance:  0.00,
	}

	store := userstore.NewStore()
	if err = store.CreateUser(c, user); err != nil {
		c.Set("Content-type", "text/html")
		return index.HandleSignUP(err.Error(), c).Render(c.UserContext(), c.Response().BodyWriter())
	}

	val := &fiber.Cookie{
		Name: "session_id",
		// Set expiry date to the past
		Expires:  time.Now().Add(-(time.Hour * 2)),
		HTTPOnly: true,
		Value:    "super_test",
		SameSite: "lax",
	}
	c.Cookie(val)
	defer sess.Save()
	sess.Set(val.Name, val.Value)

	return c.Redirect("/personal/dash")
}

// Hanlde Get request
func HandleSignUp(c *fiber.Ctx) error {

	cookie := c.Cookies("session_id")
	if len(cookie) > 1 {
		return c.Redirect("/personal/dash")
	}

	c.Set("Content-type", "text/html")
	return index.HandleSignUP("", c).Render(c.UserContext(), c.Response().BodyWriter())
}

func CheckAuth(c *fiber.Ctx) error {
	if session := c.Cookies("session_id"); len(session) < 1 {
		c.ClearCookie("session_id")
	}
	c.Redirect("/", fiber.StatusNotAcceptable)
	return c.Next()
}

func HandleDashboard(c *fiber.Ctx) error {

	cookie := c.Cookies("session_id")
	if len(cookie) < 1 {
		return c.Redirect("/")
	}

	c.Set("Content-type", "text/html")
	return index.HandleDash().Render(c.UserContext(), c.Response().BodyWriter())
}

func HandleWallets(c *fiber.Ctx) error {
	return c.Redirect("/personal/kyc")
}
