// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.707
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "github.com/gofiber/fiber/v2"

var status = ""

func change(status string) string {
	status = "nav-link my-active"
	return status
}

func Cookies(c *fiber.Ctx) bool {
	cookie := c.Cookies("session_id")
	if len(cookie) > 1 {
		return false
	}
	return true
}

func Head() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<head><!-- Title --><title>Paylary  - Multi The WEB 3.0 Financial ecosystem free account to everyone</title><!-- Required Meta Tags Always Come First --><meta charset=\"utf-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1, shrink-to-fit=no\"><!-- Favicon --><link rel=\"shortcut icon\" href=\"./static/assets/favicon.ico\"><!-- Font --><link href=\"https://fonts.googleapis.com/css2?family=Open+Sans:wght@400;600&amp;display=swap\" rel=\"stylesheet\"><!-- CSS Implementing Plugins --><link rel=\"stylesheet\" href=\"./static/assets/vendor/fontawesome/css/all.min.css\"><link rel=\"stylesheet\" href=\"./static/assets/vendor/hs-mega-menu/dist/hs-mega-menu.min.css\"><script src=\"https://unpkg.com/htmx.org@2.0.1\"></script><!-- CSS Implementing Plugins --><link rel=\"stylesheet\" href=\"./static/assets/vendor/fontawesome/css/all.min.css\"><link rel=\"stylesheet\" href=\"./static/assets/vendor/@fancyapps/fancybox/dist/jquery.fancybox.min.css\"><link rel=\"stylesheet\" href=\"./static/assets/vendor/aos/dist/aos.css\"><link rel=\"stylesheet\" href=\"./static/assets/vendor/slick-carousel/slick/slick.css\"><!-- CSS Front Template --><link rel=\"stylesheet\" href=\"./static/assets/css/theme.min.css\"><link rel=\"stylesheet\" href=\"./static/assets/css/custom.css\"></head>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func Header(category string, url string, c *fiber.Ctx) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var2 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var2 == nil {
			templ_7745c5c3_Var2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<header id=\"headerWithBSDropdown\" class=\"header header-sticky-top nonpadding\"><div class=\"header-section bg-white nonpadding\"><div id=\"logoAndNavWithBSDropdown\" class=\"container nonpadding\"><nav class=\"js-mega-menu navbar navbar-expand-lg hs-menu-initialized hs-menu-horizontal nonpadding\"><div class=\"navbar-brand-wrapper\"><!-- Logo --><a class=\"navbar-brand nonpadding logo\" href=\"/\"><svg width=\"45%\" height=\"100%\" viewBox=\"0 0 65 64\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M32.5 64C54.5914 64 64.5 54.0914 64.5 32C64.5 9.90862 54.5914 0 32.5 0C10.4086 0 0.5 9.90862 0.5 32C0.5 54.0914 10.4086 64 32.5 64ZM35.3857 13.5696H26.4834L24.9217 18.292H29.8662V28.2395H35.3857V13.5696ZM42.6504 40.2785C42.6504 45.8856 38.105 50.4311 32.4979 50.4311C26.8908 50.4311 22.3453 45.8856 22.3453 40.2785C22.3453 34.6714 26.8908 30.126 32.4979 30.126C38.105 30.126 42.6504 34.6714 42.6504 40.2785ZM38.0413 40.2785C38.0413 43.3401 35.5594 45.822 32.4978 45.822C29.4363 45.822 26.9544 43.3401 26.9544 40.2785C26.9544 37.217 29.4363 34.735 32.4978 34.735C35.5594 34.735 38.0413 37.217 38.0413 40.2785Z\" fill=\"#16D629\"></path><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M32.5 64C54.5914 64 64.5 54.0914 64.5 32C64.5 9.90862 54.5914 0 32.5 0C10.4086 0 0.5 9.90862 0.5 32C0.5 54.0914 10.4086 64 32.5 64ZM35.3857 13.5696H26.4834L24.9217 18.292H29.8662V28.2395H35.3857V13.5696ZM42.6504 40.2785C42.6504 45.8856 38.105 50.4311 32.4979 50.4311C26.8908 50.4311 22.3453 45.8856 22.3453 40.2785C22.3453 34.6714 26.8908 30.126 32.4979 30.126C38.105 30.126 42.6504 34.6714 42.6504 40.2785ZM38.0413 40.2785C38.0413 43.3401 35.5594 45.822 32.4978 45.822C29.4363 45.822 26.9544 43.3401 26.9544 40.2785C26.9544 37.217 29.4363 34.735 32.4978 34.735C35.5594 34.735 38.0413 37.217 38.0413 40.2785Z\" fill=\"currentColor\"></path></svg></a><!-- End Logo --></div><!-- Nav --><!-- Responsive Toggle Button --><button type=\"button\" class=\"navbar-toggler btn btn-icon btn-sm rounded-circle nonpadding\" aria-label=\"Toggle navigation\" aria-expanded=\"false\" aria-controls=\"navBarWithBSDropdown\" data-toggle=\"collapse\" data-target=\"#navBarWithBSDropdown\"><span class=\"navbar-toggler-default\"><svg width=\"14\" height=\"14\" viewBox=\"0 0 18 18\" xmlns=\"http://www.w3.org/2000/svg\"><path fill=\"currentColor\" d=\"M17.4,6.2H0.6C0.3,6.2,0,5.9,0,5.5V4.1c0-0.4,0.3-0.7,0.6-0.7h16.9c0.3,0,0.6,0.3,0.6,0.7v1.4C18,5.9,17.7,6.2,17.4,6.2z M17.4,14.1H0.6c-0.3,0-0.6-0.3-0.6-0.7V12c0-0.4,0.3-0.7,0.6-0.7h16.9c0.3,0,0.6,0.3,0.6,0.7v1.4C18,13.7,17.7,14.1,17.4,14.1z\"></path></svg></span> <span class=\"navbar-toggler-toggled\"><svg width=\"14\" height=\"14\" viewBox=\"0 0 18 18\" xmlns=\"http://www.w3.org/2000/svg\"><path fill=\"currentColor\" d=\"M11.5,9.5l5-5c0.2-0.2,0.2-0.6-0.1-0.9l-1-1c-0.3-0.3-0.7-0.3-0.9-0.1l-5,5l-5-5C4.3,2.3,3.9,2.4,3.6,2.6l-1,1 C2.4,3.9,2.3,4.3,2.5,4.5l5,5l-5,5c-0.2,0.2-0.2,0.6,0.1,0.9l1,1c0.3,0.3,0.7,0.3,0.9,0.1l5-5l5,5c0.2,0.2,0.6,0.2,0.9-0.1l1-1 c0.3-0.3,0.3-0.7,0.1-0.9L11.5,9.5z\"></path></svg></span></button><!-- End Responsive Toggle Button --><!-- Navigation --><div id=\"navBarWithBSDropdown\" class=\"collapse navbar-collapse nonpadding\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		switch category {
		case "b2c":
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<ul class=\"nav nav-segment\" id=\"navTab3\" role=\"tablist\"><li role=\"presentation\" class=\"nav-item\"><a id=\"nav-resultTab3\" href=\"/\" data-bs-toggle=\"pill\" data-bs-target=\"#nav-result3\" role=\"tab\" aria-controls=\"nav-result3\" aria-selected=\"true\" class=\"nav-link active pl-4 pr-4\">B2C</a></li><li class=\"nav-item\" role=\"presentation\"><a id=\"nav-htmlTab3\" href=\"/b2b\" data-bs-toggle=\"pill\" data-bs-target=\"#nav-html3\" role=\"tab\" aria-controls=\"nav-html3\" aria-selected=\"false\" tabindex=\"-1\" class=\"nav-link pl-4 pr-4\">B2B</a></li></ul>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		case "b2b":
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<ul class=\"nav nav-segment\" id=\"navTab3\" role=\"tablist\"><li role=\"presentation\" class=\"nav-item\"><a id=\"nav-resultTab3\" href=\"/\" data-bs-toggle=\"pill\" data-bs-target=\"#nav-result3\" role=\"tab\" aria-controls=\"nav-result3\" aria-selected=\"true\" class=\"nav-link  pl-4 pr-4\">B2C</a></li><li class=\"nav-item\" role=\"presentation\"><a id=\"nav-htmlTab3\" href=\"/b2b\" data-bs-toggle=\"pill\" data-bs-target=\"#nav-html3\" role=\"tab\" aria-controls=\"nav-html3\" aria-selected=\"false\" tabindex=\"-1\" class=\"nav-link active pl-4 pr-4\">B2B</a></li></ul>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		switch category {
		case "b2c":
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<ul class=\"navbar-nav\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if url == "wallet" {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li class=\"navbar-nav-item\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var3 = []any{change(status)}
				templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var3...)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<a class=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var4 string
				templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var3).String())
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/views/components/components.templ`, Line: 1, Col: 0}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" href=\"/wallet\">Wallet</a></li>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li class=\"navbar-nav-item\"><a class=\"nav-link\" href=\"/wallet\">Wallet</a></li>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			if url == "exchange" {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li class=\"navbar-nav-item\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var5 = []any{change(status)}
				templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var5...)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<a class=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var6 string
				templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var5).String())
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/views/components/components.templ`, Line: 1, Col: 0}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" href=\"/exchange\">Exchange</a></li>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li class=\"navbar-nav-item\"><a class=\"nav-link\" href=\"/exchange\">Exchange</a></li>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			if url == "about" {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li class=\"navbar-nav-item\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var7 = []any{change(status)}
				templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var7...)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<a class=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var8 string
				templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var7).String())
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/views/components/components.templ`, Line: 1, Col: 0}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" href=\"/about\">About</a></li>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li class=\"navbar-nav-item\"><a class=\"nav-link\" href=\"/about\">About</a></li>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			if url == "contacts" {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li class=\"navbar-nav-item\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var9 = []any{change(status)}
				templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var9...)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<a class=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var10 string
				templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var9).String())
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/views/components/components.templ`, Line: 1, Col: 0}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" href=\"/contacts\">Contacts</a></li>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li class=\"navbar-nav-item\"><a class=\"nav-link\" href=\"/contacts\">Contacts</a></li>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			if ok := Cookies(c); ok {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li class=\"nav-item my-button\"><a class=\"btn btn-primary btn-transition\" href=\"/login\">Sign in</a></li><li class=\"nav-item my-button\"><a class=\"btn btn-transition\" href=\"/signup\">Sign up</a></li>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li class=\"nav-item my-button\"><a href=\"/personal/dash\" class=\"\" style=\"width: 100px;\"><img color=\"white\" src=\"./static/assets/svg/logos/person.svg\" width=\"30px;\"></a></li>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</ul>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		case "b2b":
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<ul class=\"navbar-nav\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if url == "payments" {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li class=\"navbar-nav-item\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var11 = []any{change(status)}
				templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var11...)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<a class=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var12 string
				templ_7745c5c3_Var12, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var11).String())
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/views/components/components.templ`, Line: 1, Col: 0}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var12))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" href=\"/payments\">Payments</a></li>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li class=\"navbar-nav-item\"><a class=\"nav-link\" href=\"/payments\">Payments</a></li>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			if url == "api" {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li class=\"navbar-nav-item\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var13 = []any{change(status)}
				templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var13...)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<a class=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var14 string
				templ_7745c5c3_Var14, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var13).String())
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/views/components/components.templ`, Line: 1, Col: 0}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var14))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" href=\"/api\">API & SDK</a></li>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li class=\"navbar-nav-item\"><a class=\"nav-link\" href=\"/api\">API & SDK</a></li>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			if url == "compl" {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li class=\"navbar-nav-item\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var15 = []any{change(status)}
				templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var15...)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<a class=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var16 string
				templ_7745c5c3_Var16, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var15).String())
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/views/components/components.templ`, Line: 1, Col: 0}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var16))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" href=\"/compliance\">Compliance</a></li>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li class=\"navbar-nav-item\"><a class=\"nav-link\" href=\"/compliance\">Compliance</a></li>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			if url == "about" {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li class=\"navbar-nav-item\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var17 = []any{change(status)}
				templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var17...)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<a class=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var18 string
				templ_7745c5c3_Var18, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var17).String())
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/views/components/components.templ`, Line: 1, Col: 0}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var18))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" href=\"/about\">About</a></li>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li class=\"navbar-nav-item\"><a class=\"nav-link\" href=\"/about\">About</a></li>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			if url == "contacts" {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li class=\"navbar-nav-item\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var19 = []any{change(status)}
				templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var19...)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<a class=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var20 string
				templ_7745c5c3_Var20, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var19).String())
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/views/components/components.templ`, Line: 1, Col: 0}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var20))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" href=\"/contacts\">Contacts</a></li>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li class=\"navbar-nav-item\"><a class=\"nav-link\" href=\"/contacts\">Contacts</a></li>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			if ok := Cookies(c); ok {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li class=\"nav-item my-button\"><a class=\"btn btn-primary btn-transition\" href=\"/login\">Sign in</a></li><li class=\"nav-item my-button\"><a class=\"btn btn-transition\" href=\"/signup\">Sign up</a></li>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li class=\"nav-item my-button\"><a href=\"/personal/dash\" class=\"\" style=\"width: 100px;\"><img color=\"white\" src=\"./static/assets/svg/logos/person.svg\" width=\"30px;\"></a></li>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</ul>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><!-- End Navigation --></nav><!-- End Nav --></div></div></header>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func BottomScripts() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var21 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var21 == nil {
			templ_7745c5c3_Var21 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!-- Go to Top --><a class=\"js-go-to go-to position-fixed\" href=\"javascript:;\" style=\"visibility: hidden;\" data-hs-go-to-options=\"{\n       &#34;offsetTop&#34;: 700,\n       &#34;position&#34;: {\n         &#34;init&#34;: {\n           &#34;right&#34;: 15\n         },\n         &#34;show&#34;: {\n           &#34;bottom&#34;: 15\n         },\n         &#34;hide&#34;: {\n           &#34;bottom&#34;: -15\n         }\n       }\n     }\"><i class=\"fas fa-angle-up\"></i></a><!-- End Go to Top --><!-- JS Global Compulsory  --><script src=\"./static/assets/vendor/jquery/dist/jquery.min.js\"></script><script src=\"./static/assets/vendor/jquery-migrate/dist/jquery-migrate.min.js\"></script><script src=\"./static/assets/vendor/bootstrap/dist/js/bootstrap.bundle.min.js\"></script><!-- JS Implementing Plugins --><script src=\"./static/assets/vendor/hs-header/dist/hs-header.min.js\"></script><script src=\"./static/assets/vendor/hs-go-to/dist/hs-go-to.min.js\"></script><script src=\"./static/assets/vendor/hs-unfold/dist/hs-unfold.min.js\"></script><script src=\"./static/assets/vendor/hs-mega-menu/dist/hs-mega-menu.min.js\"></script><script src=\"./static/assets/vendor/hs-show-animation/dist/hs-show-animation.min.js\"></script><script src=\"./static/assets/vendor/jquery-validation/dist/jquery.validate.min.js\"></script><!-- JS Front --><script src=\"./static/assets/js/theme.min.js\"></script><!-- JS Plugins Init. --><script>\n    $(document).on('ready', function () {\n      // INITIALIZATION OF HEADER\n      // =======================================================\n      var header = new HSHeader($('#header')).init();\n\n\n      // INITIALIZATION OF MEGA MENU\n      // =======================================================\n      var megaMenu = new HSMegaMenu($('.js-mega-menu'), {\n        desktop: {\n          position: 'left'\n        }\n      }).init();\n\n\n      // INITIALIZATION OF UNFOLD\n      // =======================================================\n      var unfold = new HSUnfold('.js-hs-unfold-invoker').init();\n\n\n      // INITIALIZATION OF SHOW ANIMATIONS\n      // =======================================================\n      $('.js-animation-link').each(function () {\n        var showAnimation = new HSShowAnimation($(this)).init();\n      });\n\n\n      // INITIALIZATION OF FORM VALIDATION\n      // =======================================================\n      $('.js-validate').each(function() {\n        $.HSCore.components.HSValidation.init($(this), {\n          rules: {\n            confirmPassword: {\n              equalTo: '#signupPassword'\n            }\n          }\n        });\n      });\n\n\n      // INITIALIZATION OF GO TO\n      // =======================================================\n      $('.js-go-to').each(function () {\n        var goTo = new HSGoTo($(this)).init();\n      });\n    });\n  </script><!-- IE Support --><script>\n    if (/MSIE \\d|Trident.*rv:/.test(navigator.userAgent)) document.write('<script src=\"./static/assets/vendor/babel-polyfill/dist/polyfill.js\"><\\/script>');\n  </script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func Footer() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var22 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var22 == nil {
			templ_7745c5c3_Var22 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<footer class=\"bg-dark\"><div class=\"container\"><div class=\"space-top-2 space-bottom-1 space-bottom-lg-2\"><div class=\"row justify-content-lg-between\"><div class=\"col-lg-4 ml-lg-auto mb-5 mb-lg-0\"><!-- Logo --><!-- End Logo --><!-- Nav Link --><ul class=\"nav nav-sm nav-x-0 nav-white flex-column\"><li class=\"nav-item\"><span class=\"media\"><span class=\"text-white\">Paylary <i class=\"far fa-copyright\"></i> 2024</span></span></li><li class=\"nav-item\"><span class=\"media\"><span class=\"text-white\">support@paylary.com </span></span></li></ul><!-- End Nav Link --></div><div class=\"col-6 col-md-4 col-lg mb-5 mb-lg-0\"><h5 class=\"text-white\">Company</h5><!-- Nav Link --><ul class=\"nav nav-sm nav-x-0 nav-white flex-column\"><li class=\"nav-item\"><a class=\"nav-link\" href=\"/about\">About</a></li><li class=\"nav-item\"><a class=\"nav-link\" href=\"/contacts\">Contacts</a></li><li class=\"nav-item\"><a class=\"nav-link\" href=\"/\">B2C</a></li><li class=\"nav-item\"><a class=\"nav-link\" href=\"/b2b\">B2B</a></li><li class=\"nav-item\"></li></ul><!-- End Nav Link --></div><div class=\"col-6 col-md-4 col-lg mb-5 mb-lg-0\"><h5 class=\"text-white\">Features</h5><!-- Nav Link --><ul class=\"nav nav-sm nav-x-0 nav-white flex-column\"><li class=\"nav-item\"><a class=\"nav-link\" href=\"/wallet\">Wallet</a></li><li class=\"nav-item\"><a class=\"nav-link\" href=\"/exchange\">Exchange</a></li></ul><!-- End Nav Link --></div><div class=\"col-6 col-md-4 col-lg\"><h5 class=\"text-white\">Resources</h5><!-- Nav Link --><ul class=\"nav nav-sm nav-x-0 nav-white flex-column\"><li class=\"nav-item\"><a class=\"nav-link\" href=\"/contacts\">Support</a></li><li class=\"nav-item\"><a class=\"nav-link\" href=\"/security\">Security</a></li></ul><!-- End Nav Link --></div></div></div><hr class=\"opacity-xs my-0\"><div class=\"space-1\"><div class=\"row align-items-md-center mb-7\"><div class=\"col-md-6 mb-4 mb-md-0\"><!-- Nav Link --><ul class=\"nav nav-sm nav-white nav-x-sm align-items-center\"><li class=\"nav-item\"><a class=\"nav-link\" href=\"/privacy\">Privacy &amp; Policy</a></li><li class=\"nav-item opacity mx-3\">/</li><li class=\"nav-item\"><a class=\"nav-link\" href=\"/terms\">Terms &amp; Conditions</a></li><li class=\"nav-item opacity mx-3\">/</li><li class=\"nav-item\"><a class=\"nav-link\" href=\"/fraud\">Anti-fraud Policy</a></li></ul><!-- End Nav Link --></div><div class=\"col-md-6 text-md-right\"><ul class=\"list-inline mb-0\"><!-- Social Networks --><li class=\"list-inline-item\"><a class=\"btn btn-xs btn-icon btn-soft-light\" href=\"#\"><i class=\"fab fa-medium\"></i></a></li><li class=\"list-inline-item\"><a class=\"btn btn-xs btn-icon btn-soft-light\" href=\"#\"><i class=\"fab fa-discord\"></i></a></li><li class=\"list-inline-item\"><a class=\"btn btn-xs btn-icon btn-soft-light\" href=\"#\"><i class=\"fab fa-twitter\"></i></a></li><li class=\"list-inline-item\"><a class=\"btn btn-xs btn-icon btn-soft-light\" href=\"#\"><i class=\"fab fa-github\"></i></a></li><!-- End Social Networks --></ul></div></div></div></div></footer>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
