package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("validateArguments", func() {
	Context("when no arguments are provided", func() {
		It("returns false", func() {
			Expect(validateArguments([]string{})).To(BeFalse())
		})
	})

	Context("when a single letter is provided", func() {
		It("returns true", func() {
			Expect(validateArguments([]string{"A"})).To(BeTrue())
		})
	})

	Context("when multiple arguments are provided", func() {
		It("returns false", func() {
			Expect(validateArguments([]string{"A", "B"})).To(BeFalse())
		})
	})

	Context("when multiple letters are provided", func() {
		It("returns false", func() {
			Expect(validateArguments([]string{"AB"})).To(BeFalse())
		})
	})

	Context("when a single non-letter character is provided", func() {
		It("returns false", func() {
			Expect(validateArguments([]string{"1"})).To(BeFalse())
		})
	})
})
