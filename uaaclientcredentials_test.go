package uaaclientcredentials

import (
	"net/url"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
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

	Describe("Bearer Tokens", func() {
		BeforeEach(func() {
			uaaCC, _ = New(url, false, "client_id", "client_secret")
		})

		It("should return a properly formatted bearer token", func() {
			token := uaaCC.GetBearerToken()
			Expect(token).NotTo(BeNil())
			Expect(token).NotTo(Equal(""))
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
})
