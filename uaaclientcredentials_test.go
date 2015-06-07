package uaaclientcredentials

import (
	"net/url"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
)

var _ = Describe("Uaaclientcredentials", func() {
	var url *url.URL
	var uaaCC *UaaClientCredentials

	BeforeEach(func() {
		url, _ = url.Parse("https://uaa.at.your.place.com")
	})

	Describe("Creationism", func() {
		It("makes an initiliazed object", func() {
			uaaCC, _ := New(url, true, "client_id", "client_secret")
			Expect(uaaCC).NotTo(BeNil())
		})

		It("should complain about an empty client id", func() {
			uaaCC, err := New(url, false, "", "client_secret")
			Expect(uaaCC).To(BeNil())
			Expect(err).ToNot(BeNil())
		})

		It("should complain about an empty client secret", func() {
			uaaCC, err := New(url, false, "client_id", "")
			Expect(uaaCC).To(BeNil())
			Expect(err).ToNot(BeNil())
		})

	})

	Describe("SSL Validation", func() {

		It("Should skip ssl validation", func() {
			uaaCC, _ = New(url, true, "client_id", "client_secret")
			config := uaaCC.getTLSConfig()
			Expect(config.InsecureSkipVerify).To(BeTrue())
		})

		It("Should skip not ssl validation", func() {
			uaaCC, _ = New(url, false, "client_id", "client_secret")
			config := uaaCC.getTLSConfig()
			Expect(config.InsecureSkipVerify).To(BeFalse())
		})
	})

	Describe("Token Acquisition", func() {

		var server *ghttp.Server
		BeforeEach(func() {
			server = ghttp.NewTLSServer()
			url, _ = url.Parse(server.URL())
			server.AppendHandlers(
				ghttp.VerifyRequest("GET", "/aouth/token"),
				ghttp.VerifyBasicAuth("client_id", "client_secret"),
			)
			uaaCC, _ = New(url, false, "client_id", "client_secret")
		})

		AfterEach(func() {
			server.Close()
		})

		It("should fetch a credential", func() {
			uaaCC.getToken()
			Ω(server.ReceivedRequests()).Should(HaveLen(1))
			Expect(uaaCC.authorizationToken).NotTo(BeNil())
		})

		It("should ask for client credentials", func() {
		})

		It("should use the client id & secret for basic auth", func() {

		})
	})

	Describe("Bearer Tokens", func() {
		BeforeEach(func() {
			uaaCC, _ = New(url, true, "client_id", "client_secret")
			uaaCC.authorizationToken = "test_token"
		})

		It("should return a properly formatted bearer token", func() {
			token := uaaCC.GetBearerToken()
			Expect(token).To(Equal("bearer test_token"))
		})
	})
})
