package toggle_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/xchapter7x/toggle"
)

var controlSuccessStatus string = "true"

func successGetenvMock(fs string) (status string) {
	status = controlSuccessStatus
	return
}

func failureGetenvMock(fs string) (status string) {
	status = ""
	return
}

var _ = Describe("toggle package", func() {
	Describe("defaultEngine struct", func() {
		Describe("GetFeatureStatusValue function", func() {
			It("Should return the result of getenv and have nil error on success", func() {
				engine := &toggle.DefaultEngine{
					Getenv: successGetenvMock,
				}
				res, err := engine.GetFeatureStatusValue("")
				Expect(res).To(Equal(controlSuccessStatus))
				Ω(err).Should(BeNil())
			})

			It("Should return non nil err on failed call", func() {
				engine := &toggle.DefaultEngine{
					Getenv: failureGetenvMock,
				}
				_, err := engine.GetFeatureStatusValue("")
				Ω(err).ShouldNot(BeNil())
			})
		})
	})
})
