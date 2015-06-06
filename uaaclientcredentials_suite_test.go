package uaaclientcredentials_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestUaaclientcredentials(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Uaaclientcredentials Suite")
}
