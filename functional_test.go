package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Application", func() {
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

type StringOutputWriter struct {
	Output string
}

func (w *StringOutputWriter) WriteString(s string) (n int, err error) {
	w.Output += s

	return len(s), nil
}
