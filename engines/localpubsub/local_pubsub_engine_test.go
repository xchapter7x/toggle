package localpubsub_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/xchapter7x/toggle/engines/localengine"
	"github.com/xchapter7x/toggle/engines/localpubsub"
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

var _ = Describe("localpubsub package", func() {
	Describe("localpubsub struct", func() {
		Describe("GetFeatureStatusValue function", func() {
			localEngineSuccessMock := &localengine.LocalEngine{
				Getenv: successGetenvMock,
			}

			localEngineFailureMock := &localengine.LocalEngine{
				Getenv: failureGetenvMock,
			}

			It("Should return the result of getenv and have nil error on success", func() {
				engine := &localpubsub.LocalPubSubEngine{
					LocalEngine: localEngineSuccessMock,
				}
				res, err := engine.GetFeatureStatusValue("")
				Expect(res).To(Equal(controlSuccessStatus))
				Ω(err).Should(BeNil())
			})

			It("Should return non nil err on failed call", func() {
				engine := &localpubsub.LocalPubSubEngine{
					LocalEngine: localEngineFailureMock,
				}
				_, err := engine.GetFeatureStatusValue("")
				Ω(err).ShouldNot(BeNil())
			})
		})
	})
})
