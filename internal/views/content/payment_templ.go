// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.707
package content

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func PayementContent() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<main id=\"content\" role=\"main\"><!-- Hero --><div class=\"d-lg-flex position-relative\"><div class=\"container d-lg-flex align-items-lg-center content-space-t-3 content-space-lg-0 min-vh-lg-100\"><div class=\"row justify-content-lg-between align-items-lg-center mb-10\"><div class=\"col-md-6 col-lg-5\"><!-- Heading --><div class=\"mb-5\"><h1>Build amazing merchant services with digital asset payments</h1><p>Don’t be bound by traditional 9-to-5 banking systems, manual merchant settlement and payouts. The Payments Engine is a suite of tools to accept, orchestrate and settle digital asset payments across any blockchain and geography.</p></div><!-- End Title & Description --></div><!-- End Col --><div class=\"col-sm-7 col-md-6 d-none d-md-block space-top-2\"><img class=\"img-fluid\" src=\"./static/assets/img/others/payment.jpg\" alt=\"Image Description\" style=\"border-radius:30px;\"></div><!-- End Col --></div><!-- Heading --><!-- End Title & Description --><!-- SVG Shape --><!-- End SVG Shape --></div></div><!-- End Hero --><!-- Card Grid --><div class=\"container content-space-t-2 content-space-b-2 content-space-b-lg-3 space-top-3 space-bottom-3\"><div class=\"row align-items-md-center\"><div class=\"col-md-7 mb-md-0\"><div class=\"mx-sm-auto aos-init aos-animate\" data-aos=\"fade-up\"><!-- Card --><div class=\"card bg-soft-primary text-center\"><div class=\"card-body\"><div class=\"w-75 mx-auto\"></div></div><div class=\"w-75 mx-auto\"><img class=\"img-fluid rounded-top-2\" src=\"./static/assets/img/others/b2b_flow.png\" alt=\"Image Description\"></div></div><!-- End Card --></div></div><!-- End Col --><div class=\"col-md-5 order-md-2\"><!-- Heading --><div class=\"mb-5\"><h2 class=\"mb-3\">Customize your digital asset payment flows</h2><p>Paylary provides a connectivity layer to on/off ramps, exchanges, and banks so you can build dynamic flows that Convert, Transfer, and Disburse digital asset payments with a single solution.</p></div><!-- End Heading --><!-- List Checked --><ul class=\"list-checked list-checked-soft-bg-primary list-checked-lg\"><li class=\"list-checked-item\">Instant integration that enabless support of fiat to stablecoin conversion</li><li class=\"list-checked-item\">Automate digital asset transfers</li><li class=\"list-checked-item\">Disburse payments with a single flow</li></ul><!-- End List Checked --></div><!-- End Col --></div><!-- End Row --></div><!-- End Card Grid --><!-- Features --><div class=\"overflow-hidden\"><div class=\"container space-2\"><div class=\"text-center mx-md-auto mb-5 mb-md-9\"><span class=\"d-block font-weight-bold text-cap mb-2\">Paylary</span><h2>Visibility at every step of the payment flow</h2></div><div class=\"row justify-content-lg-between align-items-lg-center\"><div class=\"col-lg-6 mb-9 mb-lg-0\"><!-- Mockups --><div class=\"\"><div class=\"\"><img class=\"device-iphone-x-frame\" src=\"./static/assets/img/others/b2b_dash.jpg\" alt=\"Image Description\" style=\"border-radius:30px;/* border: 5px solid #ffffff; */\"></div></div><!-- End Mockups --></div><div class=\"col-lg-5\"><div class=\"mb-5\"></div><!-- Icon Block --><ul class=\"step step-dashed mb-7\"><li class=\"step-item\"><div class=\"step-content-wrapper\"><div class=\"step-content\"><h3 class=\"h4\">Expand to new merchant markets</h3><p>With our offering, customers can launch new blockchain-based payment services that offer instant merchant settlement on stablecoin 24/7/365, and benefit from additional revenue, stronger capital efficiency, and the ability to serve customers in more geographies.</p></div></div></li><li class=\"step-item\"><div class=\"step-content-wrapper\"><div class=\"step-content\"><h4>Move money across borders in minutes</h4><p class=\"mb-0\">Provide merchants with more options to transfer funds cross border with digital asset settlements in minutes, not days. Only Fireblocks enables you to connect to exchanges and on/off ramps so you can orchestrate stablecoin settlement flows across corridors with a single click.</p></div></div></li><li class=\"step-item\"><div class=\"step-content-wrapper\"><div class=\"step-content\"><h4>Predictable digital asset settlements, constant visibility</h4><p class=\"mb-0\">Enable merchants and your treasury teams to operate with complete predictability of settlement and visibility into the entire digital asset payments flow. Fireblocks provides a real-time view into the exact location of funds, enhancing capital and liquidity management.</p></div></div></li></ul><!-- End Icon Block --></div></div></div></div><!-- End Features --><!-- Testimonials --><div class=\"container content-space-2 content-space-lg-3 space-top-3\"><!-- End Row --><!-- FAQ --><!-- End FAQ --></div><!-- End Testimonials --><!-- Stats --><!-- End Stats --><!-- Card Grid --><!-- End Card Grid --></main>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
