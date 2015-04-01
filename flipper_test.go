package toggle_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/xchapter7x/toggle"
)

func randomA(a, b int) int {
	return a * b
}
func randomB(a, b int) int {
	return a + b
}

var _ = Describe("Tgl", func() {
	Context("calling the full chain", func() {
		var (
			tgl     *Tgl
			control int
			argA    int
			argB    int
			rval    int
		)

		BeforeEach(func() {
			tgl = NewTgl()
			control = 0
			argA = 5
			argB = 2
			rval = control

			Ω(func() {
				_, err := tgl.Flag("randomFeature").
					On(randomA).
					Off(randomB).
					Args(argA, argB).
					Returns(&rval).
					Run()
				Ω(err).Should(BeNil())
			}).ShouldNot(Panic())
		})

		It("should return the proper value to the return variable pointer given", func() {
			controlRes := randomB(argA, argB)
			Ω(rval).ShouldNot(Equal(control))
			Ω(rval).Should(Equal(controlRes))
		})
	})
})
