// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.707
package content

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Wallet() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<main id=\"content\" role=\"main\"><!-- Hero Section --><!-- End Hero Section --><!--  Section 2--><!-- end Section 2--><!--  Section 3--><div class=\"container space-2 space-top-lg-3 mobile\"><div class=\"row justify-content-lg-between align-items-lg-center\"><div class=\"col-lg-5 mb-9 mb-lg-0\"><div class=\"mb-5\"><h2>True crypto ownership. Powerful Web3 experiences</h2><p>Unlock the power of your cryptocurrency assets and explore the world of Web3 with Paylary.</p></div><div class=\"w-md-50 w-lg-80 mb-2\"><!-- Fancybox --><div class=\"text-center rounded-lg py-11 px-5\"><h4 class=\"mb-3\">Get our mobile app</h4><a class=\"btn btn-icon btn-indigo rounded-circle mr-2\" href=\"#\"><i class=\"fab fa-apple\"></i></a> <a class=\"btn btn-icon btn-indigo rounded-circle\" href=\"#\"><i class=\"fab fa-google-play\"></i></a></div><!-- End Fancybox --></div></div><div class=\"col-lg-6\"><div class=\"position-relative mx-auto aos-init aos-animate\" data-aos=\"fade-up\"><img class=\"device-iphone-x-frame\" src=\"./static/assets/img/others/wallet.png\" alt=\"Image Description\"><!-- Device Mockup --><!-- End Device Mockup --><!-- SVG Component --><div class=\"position-absolute bottom-0 right-0 max-w-50rem w-100 z-index-n1 mx-auto\"></div><!-- End SVG Component --></div></div></div></div><!-- End Features Section --><!-- Features Section --><div class=\"container space-top-lg-3 space-bottom-lg-2\"><div class=\"row justify-content-lg-between align-items-lg-center\"><div class=\"col-lg-5 order-lg-2 mb-9 mb-lg-0\"><div class=\"mb-5\"><h2>Simple. Seamless.</h2><p>Enjoy a smooth mobile app and desktop experience with easy-to-use, powerful tools to support your entire Web3 journey. Take control of your crypto. Avoid complicated steps and deposit directly to your wallet from other exchanges.</p></div></div><div class=\"col-lg-6 order-lg-1\"><div class=\"position-relative mx-auto aos-init aos-animate\" data-aos=\"fade-up\"><!-- Device Mockup --><div class=\"\"><img class=\"device-iphone-x-frame\" src=\"./static/assets/img/svg/wallet2.svg\" alt=\"Image Description\"></div><!-- End Device Mockup --><!-- SVG Component --><div class=\"position-absolute bottom-0 right-0 max-w-50rem w-100 z-index-n1 mx-auto\"></div><!-- End SVG Component --></div></div></div></div><!-- End Features Section --><!-- Features Section --><div class=\"container space-top-2 space-top-lg-3 space-bottom-lg-2\"><div class=\"row justify-content-lg-between align-items-lg-center\"><div class=\"col-lg-5 mb-9 mb-lg-0\"><div class=\"mb-4\"><h2>Stay private and secure</h2><p>Rest easy knowing that our privacy and security measures keep you in control of your data and digital assets, while also keeping them safe.</p></div><div class=\"media\"><span class=\"icon icon-xs icon-indigo icon-circle mt-1 mr-3\"><i class=\"fas fa-check fa-xs\"></i></span><div class=\"media-body\"><p>Secure multi- currency wallet</p></div></div><div class=\"media\"><span class=\"icon icon-xs icon-indigo icon-circle mt-1 mr-3\"><i class=\"fas fa-check fa-xs\"></i></span><div class=\"media-body\"><p>All features in one wallet</p></div></div><div class=\"media\"><span class=\"icon icon-xs icon-indigo icon-circle mt-1 mr-3\"><i class=\"fas fa-check fa-xs\"></i></span><div class=\"media-body\"><p>Buy, send, spend, cash out - use all the necessary functions to work with your funds in one app.</p></div></div></div><div class=\"col-lg-6 text-center\"><div class=\"position-relative mx-auto aos-init aos-animate\" data-aos=\"fade-up\"><!-- Device Mockup --><div class=\"\"><img class=\"device-iphone-x-frame\" src=\"./static/assets/img/svg/hold.svg\" alt=\"Image Description\"></div><!-- End Device Mockup --><!-- SVG Component --><div class=\"position-absolute bottom-0 right-0 max-w-50rem w-100 z-index-n1 mx-auto\"></div><!-- End SVG Component --></div></div></div></div><div class=\"container content-space-2 content-space-t-xl-3 content-space-b-lg-3\"><!-- Heading --><div class=\"w-md-75 w-lg-50 text-center mx-md-auto mb-5\"><h2>True ownership of your crypto assets</h2></div><!-- End Heading --><div class=\"text-center mb-10\"><!-- List Checked --><ul class=\"list-inline list-checked list-checked-primary\"><li class=\"list-inline-item list-checked-item\">MPC /</li><li class=\"list-inline-item list-checked-item\">Non-castodial /</li><li class=\"list-inline-item list-checked-item\">All features</li></ul><!-- End List Checked --></div><div class=\"row mb-5 mb-md-0 space-bottom-3\"><div class=\"col-sm-6 col-lg-4 mb-4 mb-lg-0\"><!-- Card --><div class=\"card card-sm h-100\"><div class=\"p-2 col-6 text-center\"><img class=\"card-img ml-9\" src=\"./static/assets/img/svg/cloud.svg\" alt=\"Image Description\" width=\"10%\"></div><div class=\"card-body\"><h4 class=\"card-title\">Added security with encryption</h4><p class=\"card-text\">Use our Encrypted Cloud Backup for increased wallet security.</p><!-- List Pointer --><!-- End List Pointer --></div></div><!-- End Card --></div><!-- End Col --><div class=\"col-sm-6 col-lg-4 mb-4 mb-lg-0\"><!-- Card --><div class=\"card card-sm h-100\"><div class=\"p-2 col-6\"><img class=\"card-img ml-8\" src=\"./static/assets/img/svg/shield.svg\" alt=\"Image Description\"></div><div class=\"card-body\"><h4 class=\"card-title\">Zero personal tracking</h4><p class=\"card-text\">We don't track any personal information, including your IP address or balances.</p><!-- List Pointer --><!-- End List Pointer --></div></div><!-- End Card --></div><!-- End Col --><div class=\"col-sm-6 col-lg-4\"><!-- Card --><div class=\"card card-sm h-100\"><div class=\"p-2 col-6\"><img class=\"card-img ml-10\" src=\"./static/assets/img/svg/flag.svg\" alt=\"Image Description\"></div><div class=\"card-body\"><h4 class=\"card-title\">Proactive alerts for risky transactions</h4><p class=\"card-text\">Stay safe with alerts for risky address and dApp connections.</p><!-- List Pointer --><!-- End List Pointer --></div></div><!-- End Card --></div><!-- End Col --></div><!-- End Row --></div><!-- End Features Section --><!-- Section 9 --><!-- End Section 9--><!-- last Section --><!-- End last Section --></main>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
