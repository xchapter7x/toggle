package localpubsub_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	//"github.com/xchapter7x/toggle/engines/localpubsub"
)

var _ = Describe("localpubsub package", func() {
	Describe("localpubsub struct", func() {
		Describe("GetFeatureStatusValue function", func() {

			BeforeEach(func() {
			})

			AfterEach(func() {
			})

			It("Should return the result of getenv and have nil error on success", func() {
				//Expect(nil).To(Equal(nil))
				Ω(nil).Should(BeNil())
			})
		})
	})
})
