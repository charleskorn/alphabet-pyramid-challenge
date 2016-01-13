package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Alphabet Pyramid Suite")
}

var _ = Describe("The application", func() {
	It("does something", func() {

	})
})
