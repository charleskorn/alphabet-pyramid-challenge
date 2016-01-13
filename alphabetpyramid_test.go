package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Alphabet Pyramid Suite")
}

type StringOutputWriter struct {
	Output string
}

func (w *StringOutputWriter) WriteString(s string) (n int, err error) {
	w.Output += s

	return len(s), nil
}

var _ = Describe("run", func() {
	runApplication := func(args []string) string {
		outputWriter := &StringOutputWriter{}
		run(args, outputWriter)
		return outputWriter.Output
	}

	Context("when no arguments are provided", func() {
		It("prints 'INVALID INPUT'", func() {
			output := runApplication([]string{})

			Expect(output).To(Equal("INVALID INPUT\n"))
		})
	})

	Context("when an invalid argument is provided", func() {
		It("prints 'INVALID INPUT'", func() {
			output := runApplication([]string{"ABCD"})

			Expect(output).To(Equal("INVALID INPUT\n"))
		})
	})

	Context("when a valid argument is provided", func() {
		It("prints nothing", func() {
			output := runApplication([]string{"C"})

			Expect(output).To(Equal(""))
		})
	})
})

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
