package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("parseArguments", func() {
	Context("when no arguments are provided", func() {
		It("returns invalid", func() {
			_, valid := parseArguments([]string{})
			Expect(valid).To(BeFalse())
		})
	})

	Context("when a single uppercase letter is provided", func() {
		It("returns that letter", func() {
			letter, valid := parseArguments([]string{"A"})
			Expect(valid).To(BeTrue())
			Expect(letter).To(Equal('A'))
		})
	})

	Context("when a single lowercase letter is provided", func() {
		It("returns the uppercase version of that letter", func() {
			letter, valid := parseArguments([]string{"d"})
			Expect(valid).To(BeTrue())
			Expect(letter).To(Equal('D'))
		})
	})

	Context("when multiple arguments are provided", func() {
		It("returns invalid", func() {
			_, valid := parseArguments([]string{"A", "B"})
			Expect(valid).To(BeFalse())
		})
	})

	Context("when multiple letters are provided", func() {
		It("returns invalid", func() {
			_, valid := parseArguments([]string{"AB"})
			Expect(valid).To(BeFalse())
		})
	})

	Context("when a single non-letter character is provided", func() {
		It("returns invalid", func() {
			_, valid := parseArguments([]string{"1"})
			Expect(valid).To(BeFalse())
		})
	})
})

var _ = Describe("generatePyramid", func() {
	Context("when the letter is 'A'", func() {
		It("returns the pyramid for the letter A", func() {
			Expect(generatePyramid('A')).To(Equal("A\n"))
		})
	})

	Context("when the letter is 'B'", func() {
		It("returns the pyramid for the letter B", func() {
			expected := " A\n" +
				"B B\n" +
				" A\n"
			Expect(generatePyramid('B')).To(Equal(expected))
		})
	})

	Context("when the letter is 'C'", func() {
		It("returns the pyramid for the letter C", func() {
			expected := "  A\n" +
				" B B\n" +
				"C   C\n" +
				" B B\n" +
				"  A\n"
			Expect(generatePyramid('C')).To(Equal(expected))
		})
	})
})
