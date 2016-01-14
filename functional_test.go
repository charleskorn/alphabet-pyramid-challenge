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

	Context("when the letter is 'A'", func() {
		It("prints just the letter 'A'", func() {
			output := runApplication([]string{"A"})

			Expect(output).To(Equal("A\n"))
		})
	})

	Context("when the letter is 'B'", func() {
		It("prints the pyramid for the letter B", func() {
			output := runApplication([]string{"B"})
			expected := " A\n" +
				"B B\n" +
				" A\n"

			Expect(output).To(Equal(expected))
		})
	})

	Context("when the letter is 'Z'", func() {
		It("prints the pyramid for the letter Z", func() {
			output := runApplication([]string{"Z"})
			expected := "                         A\n" +
				"                        B B\n" +
				"                       C   C\n" +
				"                      D     D\n" +
				"                     E       E\n" +
				"                    F         F\n" +
				"                   G           G\n" +
				"                  H             H\n" +
				"                 I               I\n" +
				"                J                 J\n" +
				"               K                   K\n" +
				"              L                     L\n" +
				"             M                       M\n" +
				"            N                         N\n" +
				"           O                           O\n" +
				"          P                             P\n" +
				"         Q                               Q\n" +
				"        R                                 R\n" +
				"       S                                   S\n" +
				"      T                                     T\n" +
				"     U                                       U\n" +
				"    V                                         V\n" +
				"   W                                           W\n" +
				"  X                                             X\n" +
				" Y                                               Y\n" +
				"Z                                                 Z\n" +
				" Y                                               Y\n" +
				"  X                                             X\n" +
				"   W                                           W\n" +
				"    V                                         V\n" +
				"     U                                       U\n" +
				"      T                                     T\n" +
				"       S                                   S\n" +
				"        R                                 R\n" +
				"         Q                               Q\n" +
				"          P                             P\n" +
				"           O                           O\n" +
				"            N                         N\n" +
				"             M                       M\n" +
				"              L                     L\n" +
				"               K                   K\n" +
				"                J                 J\n" +
				"                 I               I\n" +
				"                  H             H\n" +
				"                   G           G\n" +
				"                    F         F\n" +
				"                     E       E\n" +
				"                      D     D\n" +
				"                       C   C\n" +
				"                        B B\n" +
				"                         A\n"

			Expect(output).To(Equal(expected))
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
