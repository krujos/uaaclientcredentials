package uaaclientcredentials_test

import (
	. "github.com/krujos/uaaclientcredentials"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Uaaclientcredentials", func() {

	Describe("Creationism", func() {
		It("makes an initiliazed object", func() {
			uaaCC, _ := New("https://uaa.at.your.place.com", "client_id", "client_secret")
			Expect(uaaCC).NotTo(BeNil())
		})
	})

	Describe("Bearer Tokens", func() {
		var uaaCC *UaaClientCredentials

		BeforeEach(func() {
			uaaCC, _ = New("https://uaa.at.your.place.com", "client_id", "client_secret")
		})

		It("should return a properly formatted bearer token", func() {
			token := uaaCC.GetBearerToken()
			Expect(token).NotTo(BeNil())
			Expect(token).NotTo(Equal(""))
		})
	})
})
