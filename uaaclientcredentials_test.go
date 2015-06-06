package uaaclientcredentials_test

import (
	"net/url"

	. "github.com/krujos/uaaclientcredentials"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Uaaclientcredentials", func() {
	var url *url.URL
	BeforeEach(func() {
		url, _ = url.Parse("https://uaa.at.your.place.com")
	})

	Describe("Creationism", func() {
		It("makes an initiliazed object", func() {
			uaaCC, _ := New(url, "client_id", "client_secret")
			Expect(uaaCC).NotTo(BeNil())
		})

		It("should complain about an empty client id", func() {
			uaaCC, err := New(url, "", "client_secret")
			Expect(uaaCC).To(BeNil())
			Expect(err).ToNot(BeNil())
		})

		It("should complain about an empty client secret", func() {
			uaaCC, err := New(url, "client_id", "")
			Expect(uaaCC).To(BeNil())
			Expect(err).ToNot(BeNil())
		})

	})

	Describe("Bearer Tokens", func() {
		var uaaCC *UaaClientCredentials

		BeforeEach(func() {
			uaaCC, _ = New(url, "client_id", "client_secret")
		})

		It("should return a properly formatted bearer token", func() {
			token := uaaCC.GetBearerToken()
			Expect(token).NotTo(BeNil())
			Expect(token).NotTo(Equal(""))
		})
	})

	Describe("Token Fetching", func() {

		It("Should skip ssl validation", func() {
			Fail("NYI")
		})
	})
})
