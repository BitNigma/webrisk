// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.707
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Header() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<header id=\"headerWithBSDropdown\" class=\"header header-sticky-top\"><div class=\"header-section bg-white\"><div id=\"logoAndNavWithBSDropdown\" class=\"container\"><!-- Nav --><nav class=\"js-mega-menu navbar navbar-expand-lg hs-menu-initialized hs-menu-horizontal\"><!-- Logo --><a class=\"navbar-brand graph\" href=\"/\" aria-label=\"Front\"><svg height=\"100%\" viewBox=\"0 0 65 64\" xmlns=\"http://www.w3.org/2000/svg\" width=\"40%\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M32.5 64C54.5914 64 64.5 54.0914 64.5 32C64.5 9.90862 54.5914 0 32.5 0C10.4086 0 0.5 9.90862 0.5 32C0.5 54.0914 10.4086 64 32.5 64ZM35.3857 13.5696H26.4834L24.9217 18.292H29.8662V28.2395H35.3857V13.5696ZM42.6504 40.2785C42.6504 45.8856 38.105 50.4311 32.4979 50.4311C26.8908 50.4311 22.3453 45.8856 22.3453 40.2785C22.3453 34.6714 26.8908 30.126 32.4979 30.126C38.105 30.126 42.6504 34.6714 42.6504 40.2785ZM38.0413 40.2785C38.0413 43.3401 35.5594 45.822 32.4978 45.822C29.4363 45.822 26.9544 43.3401 26.9544 40.2785C26.9544 37.217 29.4363 34.735 32.4978 34.735C35.5594 34.735 38.0413 37.217 38.0413 40.2785Z\" fill=\"#16D629\"></path><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M32.5 64C54.5914 64 64.5 54.0914 64.5 32C64.5 9.90862 54.5914 0 32.5 0C10.4086 0 0.5 9.90862 0.5 32C0.5 54.0914 10.4086 64 32.5 64ZM35.3857 13.5696H26.4834L24.9217 18.292H29.8662V28.2395H35.3857V13.5696ZM42.6504 40.2785C42.6504 45.8856 38.105 50.4311 32.4979 50.4311C26.8908 50.4311 22.3453 45.8856 22.3453 40.2785C22.3453 34.6714 26.8908 30.126 32.4979 30.126C38.105 30.126 42.6504 34.6714 42.6504 40.2785ZM38.0413 40.2785C38.0413 43.3401 35.5594 45.822 32.4978 45.822C29.4363 45.822 26.9544 43.3401 26.9544 40.2785C26.9544 37.217 29.4363 34.735 32.4978 34.735C35.5594 34.735 38.0413 37.217 38.0413 40.2785Z\" fill=\"currentColor\"></path></svg></a><!-- End Logo --><!-- Responsive Toggle Button --><button type=\"button\" class=\"navbar-toggler btn btn-icon btn-sm rounded-circle\" aria-label=\"Toggle navigation\" aria-expanded=\"false\" aria-controls=\"navBarWithBSDropdown\" data-toggle=\"collapse\" data-target=\"#navBarWithBSDropdown\"><span class=\"navbar-toggler-default\"><svg width=\"14\" height=\"14\" viewBox=\"0 0 18 18\" xmlns=\"http://www.w3.org/2000/svg\"><path fill=\"currentColor\" d=\"M17.4,6.2H0.6C0.3,6.2,0,5.9,0,5.5V4.1c0-0.4,0.3-0.7,0.6-0.7h16.9c0.3,0,0.6,0.3,0.6,0.7v1.4C18,5.9,17.7,6.2,17.4,6.2z M17.4,14.1H0.6c-0.3,0-0.6-0.3-0.6-0.7V12c0-0.4,0.3-0.7,0.6-0.7h16.9c0.3,0,0.6,0.3,0.6,0.7v1.4C18,13.7,17.7,14.1,17.4,14.1z\"></path></svg></span> <span class=\"navbar-toggler-toggled\"><svg width=\"14\" height=\"14\" viewBox=\"0 0 18 18\" xmlns=\"http://www.w3.org/2000/svg\"><path fill=\"currentColor\" d=\"M11.5,9.5l5-5c0.2-0.2,0.2-0.6-0.1-0.9l-1-1c-0.3-0.3-0.7-0.3-0.9-0.1l-5,5l-5-5C4.3,2.3,3.9,2.4,3.6,2.6l-1,1 C2.4,3.9,2.3,4.3,2.5,4.5l5,5l-5,5c-0.2,0.2-0.2,0.6,0.1,0.9l1,1c0.3,0.3,0.7,0.3,0.9,0.1l5-5l5,5c0.2,0.2,0.6,0.2,0.9-0.1l1-1 c0.3-0.3,0.3-0.7,0.1-0.9L11.5,9.5z\"></path></svg></span></button><!-- End Responsive Toggle Button --><!-- Navigation --><div id=\"navBarWithBSDropdown\" class=\"collapse navbar-collapse\"><ul class=\"nav nav-segment\" id=\"navTab3\" role=\"tablist\"><li role=\"presentation\" class=\"nav-item\"><a id=\"nav-resultTab3\" href=\"#nav-result3\" data-bs-toggle=\"pill\" data-bs-target=\"#nav-result3\" role=\"tab\" aria-controls=\"nav-result3\" aria-selected=\"true\" class=\"nav-link active pl-4 pr-4\">B2C</a></li><li class=\"nav-item\" role=\"presentation\"><a id=\"nav-htmlTab3\" href=\"#nav-html3\" data-bs-toggle=\"pill\" data-bs-target=\"#nav-html3\" role=\"tab\" aria-controls=\"nav-html3\" aria-selected=\"false\" tabindex=\"-1\" class=\"nav-link pl-4 pr-4\">B2B</a></li></ul><ul class=\"navbar-nav\"><li class=\"navbar-nav-item\"><a class=\"nav-link\" href=\"/messenger\">Messenger</a></li><li class=\"navbar-nav-item\"><a class=\"nav-link\" href=\"/wallet\">Wallet</a></li><li class=\"navbar-nav-item\"><a class=\"nav-link\" href=\"/exchange\">Exchange</a></li><li><a class=\"nav-link \" href=\"/blog\">Card</a></li><li class=\"navbar-nav-item\"><a class=\"nav-link \" href=\"/about\">About</a></li><li class=\"navbar-nav-item\"><a class=\"nav-link \" href=\"/contacts\">Contacts</a></li></ul></div><!-- End Navigation --></nav><!-- End Nav --></div></div></header>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func Head() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<head><!-- Title --><title>KeyHold  - Multi payment solution</title><!-- Required Meta Tags Always Come First --><meta charset=\"utf-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1, shrink-to-fit=no\"><!-- Favicon --><link rel=\"shortcut icon\" href=\"./static/assets/favicon.ico\"><!-- Font --><link href=\"https://fonts.googleapis.com/css2?family=Open+Sans:wght@400;600&amp;display=swap\" rel=\"stylesheet\"><!-- CSS Implementing Plugins --><link rel=\"stylesheet\" href=\"./static/assets/vendor/fontawesome/css/all.min.css\"><link rel=\"stylesheet\" href=\"./static/assets/vendor/hs-mega-menu/dist/hs-mega-menu.min.css\"><link rel=\"stylesheet\" href=\"./static/assets/vendor/@fancyapps/fancybox/dist/jquery.fancybox.min.css\"><link rel=\"stylesheet\" href=\"./static/assets/vendor/aos/dist/aos.css\"><link rel=\"stylesheet\" href=\"./static/assets/vendor/slick-carousel/slick/slick.css\"><!-- CSS Front Template --><link rel=\"stylesheet\" href=\"./static/assets/css/theme.min.css\"></head>")
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
		templ_7745c5c3_Var3 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var3 == nil {
			templ_7745c5c3_Var3 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!-- Go to Top --><a class=\"js-go-to go-to position-fixed animated hs-go-to-prevent-event fadeOutDown\" href=\"javascript:;\" style=\"right: 15px; bottom: -15px;\" data-hs-go-to-options=\"{\n       &#34;offsetTop&#34;: 700,\n       &#34;position&#34;: {\n         &#34;init&#34;: {\n           &#34;right&#34;: 15\n         },\n         &#34;show&#34;: {\n           &#34;bottom&#34;: 15\n         },\n         &#34;hide&#34;: {\n           &#34;bottom&#34;: -15\n         }\n       }\n     }\"><i class=\"fas fa-angle-up\"></i></a><!-- End Go to Top --><!-- JS Global Compulsory  --><script src=\"./static/assets/vendor/jquery/dist/jquery.min.js\"></script><script src=\"./static/assets/vendor/jquery-migrate/dist/jquery-migrate.min.js\"></script><script src=\"./static/assets/vendor/bootstrap/dist/js/bootstrap.bundle.min.js\"></script><!-- JS Implementing Plugins --><script src=\"./static/assets/vendor/hs-header/dist/hs-header.min.js\"></script><!-- JS Implementing Plugins --><script src=\"./static/assets/vendor/hs-nav-scroller/dist/hs-nav-scroller.min.js\"></script><script src=\"./static/assets/vendor/hs-go-to/dist/hs-go-to.min.js\"></script><script src=\"./static/assets/vendor/hs-mega-menu/dist/hs-mega-menu.min.js\"></script><script src=\"./static/assets/vendor/typed.js/lib/typed.min.js\"></script><script src=\"./static/assets/vendor/@fancyapps/fancybox/dist/jquery.fancybox.min.js\"></script><script src=\"./static/assets/vendor/aos/dist/aos.js\"></script><!-- JS Front --><script src=\"./static/assets/js/theme.min.js\"></script><!-- JS Plugins Init. --><script>\n    $(document).on('ready', function () {\n      // INITIALIZATION OF HEADER\n      // =======================================================\n      var header = new HSHeader($('#header')).init();\n\n      // INITIALIZATION OF MEGA MENU\n      // =======================================================\n      var megaMenu = new HSMegaMenu($('.js-mega-menu'), {\n        desktop: {\n          position: 'left'\n        }\n      }).init();\n\n      // INITIALIZATION OF NAV SCROLLER\n       // =======================================================\n      new HsNavScroller('.js-nav-scroller')\n\n\n      // INITIALIZATION OF TEXT ANIMATION (TYPING)\n      // =======================================================\n      var typed = $.HSCore.components.HSTyped.init(\".js-text-animation\");\n\n      // INITIALIZATION OF FANCYBOX\n      // =======================================================\n      $('.js-fancybox').each(function () {\n        var fancybox = $.HSCore.components.HSFancyBox.init($(this));\n      });\n\n      // INITIALIZATION OF AOS\n      // =======================================================\n      AOS.init({\n        duration: 650,\n        once: true\n      });\n\n      // INITIALIZATION OF GO TO\n      // =======================================================\n      $('.js-go-to').each(function () {\n        var goTo = new HSGoTo($(this)).init();\n      });\n    });\n  </script><!-- IE Support --><script>\n    if (/MSIE \\d|Trident.*rv:/.test(navigator.userAgent)) document.write('<script src=\"./static/assets/vendor/babel-polyfill/dist/polyfill.js\"><\\/script>');\n  </script><style type=\"text/css\" data-typed-js-css=\"true\">\n        .typed-cursor{\n          opacity: 1;\n        }\n        .typed-cursor.typed-cursor--blink{\n          animation: typedjsBlink 0.7s infinite;\n          -webkit-animation: typedjsBlink 0.7s infinite;\n                  animation: typedjsBlink 0.7s infinite;\n        }\n        @keyframes typedjsBlink{\n          50% { opacity: 0.0; }\n        }\n        @-webkit-keyframes typedjsBlink{\n          0% { opacity: 1; }\n          50% { opacity: 0.0; }\n          100% { opacity: 1; }\n        }\n      </style>")
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
		templ_7745c5c3_Var4 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var4 == nil {
			templ_7745c5c3_Var4 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<footer class=\"bg-dark\"><div class=\"container\"><div class=\"space-top-2 space-bottom-1 space-bottom-lg-2\"><div class=\"row justify-content-lg-between\"><div class=\"col-lg-4 ml-lg-auto mb-5 mb-lg-0\"><!-- Logo --><!-- End Logo --><!-- Nav Link --><ul class=\"nav nav-sm nav-x-0 nav-white flex-column\"><li class=\"nav-item\"><span class=\"media\"><span class=\"text-white\">KeyHold <i class=\"far fa-copyright\"></i> 2024</span></span></li><li class=\"nav-item\"><span class=\"media\"><span class=\"text-white\">support@keyhold.com </span></span></li></ul><!-- End Nav Link --></div><div class=\"col-6 col-md-4 col-lg mb-5 mb-lg-0\"><h5 class=\"text-white\">Company</h5><!-- Nav Link --><ul class=\"nav nav-sm nav-x-0 nav-white flex-column\"><li class=\"nav-item\"><a class=\"nav-link\" href=\"/about\">About</a></li><li class=\"nav-item\"><a class=\"nav-link\" href=\"/contacts\">Contacts</a></li><li class=\"nav-item\"><a class=\"nav-link\" href=\"/contacts\">B2C</a></li><li class=\"nav-item\"><a class=\"nav-link\" href=\"/contacts\">B2B</a></li><li class=\"nav-item\"></li></ul><!-- End Nav Link --></div><div class=\"col-6 col-md-4 col-lg mb-5 mb-lg-0\"><h5 class=\"text-white\">Features</h5><!-- Nav Link --><ul class=\"nav nav-sm nav-x-0 nav-white flex-column\"><li class=\"nav-item\"><a class=\"nav-link\" href=\"/messenger\">Messenger</a></li><li class=\"nav-item\"><a class=\"nav-link\" href=\"/wallet\">Wallet</a></li><li class=\"nav-item\"><a class=\"nav-link\" href=\"/exchange\">Exchange</a></li><li class=\"nav-item\"><a class=\"nav-link\" href=\"/exchange\">Card</a></li></ul><!-- End Nav Link --></div><div class=\"col-6 col-md-4 col-lg\"><h5 class=\"text-white\">Resources</h5><!-- Nav Link --><ul class=\"nav nav-sm nav-x-0 nav-white flex-column\"><li class=\"nav-item\"><a class=\"nav-link\" href=\"/contacts\">Support</a></li><li class=\"nav-item\"><a class=\"nav-link\" href=\"/help-center\">Terms of Use</a></li><li class=\"nav-item\"><a class=\"nav-link\" href=\"/help-center\">AML/KYC Policy</a></li></ul><!-- End Nav Link --></div></div></div><hr class=\"opacity-xs my-0\"><div class=\"space-1\"><div class=\"row align-items-md-center mb-7\"><div class=\"col-md-6 mb-4 mb-md-0\"><!-- Nav Link --><ul class=\"nav nav-sm nav-white nav-x-sm align-items-center\"><li class=\"nav-item\"><a class=\"nav-link\" href=\"/privacy\">Privacy &amp; Policy</a></li><li class=\"nav-item opacity mx-3\">/</li><li class=\"nav-item\"><a class=\"nav-link\" href=\"/terms\">Terms &amp; Conditions</a></li><li class=\"nav-item opacity mx-3\">/</li><li class=\"nav-item\"><a class=\"nav-link\" href=\"/tech\">Anti-fraud Policy</a></li></ul><!-- End Nav Link --></div><div class=\"col-md-6 text-md-right\"><ul class=\"list-inline mb-0\"><!-- Social Networks --><li class=\"list-inline-item\"><a class=\"btn btn-xs btn-icon btn-soft-light\" href=\"#\"><i class=\"fab fa-medium\"></i></a></li><li class=\"list-inline-item\"><a class=\"btn btn-xs btn-icon btn-soft-light\" href=\"#\"><i class=\"fab fa-discord\"></i></a></li><li class=\"list-inline-item\"><a class=\"btn btn-xs btn-icon btn-soft-light\" href=\"#\"><i class=\"fab fa-twitter\"></i></a></li><li class=\"list-inline-item\"><a class=\"btn btn-xs btn-icon btn-soft-light\" href=\"#\"><i class=\"fab fa-github\"></i></a></li><!-- End Social Networks --></ul></div></div></div></div></footer>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
