// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.707
package content

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func SecurityPolicy() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<main id=\"content\" role=\"main\"><!-- Content --><div class=\"container content-space-3 content-space-lg-4\"><!-- Card --><div class=\"card card-lg space-top-3\"><!-- Header --><div class=\"card-header bg-dark py-sm-10 space-top-3\"><h1 class=\"card-title h2 text-white\">Security</h1><p class=\"card-text text-white\">Last modified: March 27, 2024 </p></div><!-- End Header --><!-- Card Body --><div class=\"card-body\"><div class=\"mb-7\"><h4>Security First, Always</h4><p>As the world’s leading cryptocurrency platform for licenses,  registrations, and security certifications, our commitment to our  customers is built on trust. We believe that security and data privacy  are the foundations of achieving mainstream cryptocurrency  adoption.</p></div><div class=\"mb-7\"><h4>Our Philosophy - Security and Privacy by Design and By Default</h4><p>We drive a Zero Trust, Defence in Depth security strategy across our systems and platforms. Data privacy assessments are built into our processes to safeguard your personal information.</p><h4>Empowering a Growth Mindset</h4><p>To continually strengthen our security posture, we invest heavily in ongoing security and privacy awareness training for all staff.</p></div><div class=\"mb-7\"><h4>100% of user cryptocurrency assets are safely held and fully backed 1:1.</h4><p>We hold all customer assets deposited on our platform in institutional-grade reserve accounts on a 1:1 basis, meaning funds are responsibly backed by Paylary and accessible at customers’ convenience. Our users can verify our reserves and their funds through our Proof of Reserves verification page, conducted by an independent third-party.</p><p>Security is baked into our coding lifecycle. Our software is peer-reviewed and uses a combination of static and dynamic source code analysis tools.</p></div><div class=\"mb-7\"><h4>Crypto Protocols</h4><h5 class=\"small\">XEdDSA and VXEdDSA</h5><p>This protocol usuing for creation and verify EdDSA-compatible signatures using public key and private key formats initially defined for the X25519 and X448 elliptic curve Diffie-Hellman functions. This document also describes \"VXEdDSA\" which extends XEdDSA to make it a verifiable random function, or VRF.</p><div class=\"mb-7\"><h5 class=\"small\">PQXDH</h5><p>This protocol is the \"PQXDH\" (or \"Post-Quantum Extended Diffie-Hellman\") key agreement protocol. PQXDH establishes a shared secret key between two parties who mutually authenticate each other based on public keys. PQXDH provides post-quantum forward secrecy and a form of cryptographic deniability but still relies on the hardness of the discrete log problem for mutual authentication in this revision of the protocol.</p></div></div><div class=\"mb-7\"><h5 class=\"small\">X3DH</h5><p>X3DH establishes a shared secret key between two parties who mutually authenticate each other based on public keys. X3DH provides forward secrecy and cryptographic deniability.</p><div class=\"mb-7\"><h5 class=\"small\">Double Ratchet</h5><p>This protocol describes the Double Ratchet algorithm, which is used by two parties to exchange encrypted messages based on a shared secret key. The parties derive new keys for every Double Ratchet message so that earlier keys cannot be calculated from later ones. The parties also send Diffie-Hellman public values attached to their messages. The results of Diffie-Hellman calculations are mixed into the derived keys so that later keys cannot be calculated from earlier ones. These properties give some protection to earlier or later encrypted messages in case of a compromise of a party's keys.</p></div></div></div><!-- End Card Body --></div><!-- End Card --></div><!-- End Content --></main>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
