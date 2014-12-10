package localpubsub_test

import (
	"sync"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/xchapter7x/toggle"
	"github.com/xchapter7x/toggle/engines/localengine"
	"github.com/xchapter7x/toggle/engines/localpubsub"
)

type PubSubConnMock struct {
}

func (c PubSubConnMock) Close() (err error) {
	return
}

func (c PubSubConnMock) Subscribe(channel ...interface{}) (err error) {
	return
}

func (c PubSubConnMock) PSubscribe(channel ...interface{}) (err error) {
	return
}

func (c PubSubConnMock) Unsubscribe(channel ...interface{}) (err error) {
	return
}

func (c PubSubConnMock) PUnsubscribe(channel ...interface{}) (err error) {
	return
}

func (c PubSubConnMock) Receive() (i interface{}) {
	return
}

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
		Context("GetFeatureStatusValue function", func() {
			var localEngineFailureMock, localEngineSuccessMock *localengine.LocalEngine
			var engine *localpubsub.LocalPubSubEngine

			BeforeEach(func() {
				localEngineSuccessMock = &localengine.LocalEngine{
					Getenv: successGetenvMock,
				}

				localEngineFailureMock = &localengine.LocalEngine{
					Getenv: failureGetenvMock,
				}
			})

			AfterEach(func() {
				localEngineFailureMock.Close()
				localEngineSuccessMock.Close()
				engine.Close()
			})
			It("Should return the result of getenv and have nil error on success", func() {
				engine = &localpubsub.LocalPubSubEngine{
					LocalEngine: localEngineSuccessMock,
					PubSub:      &PubSubConnMock{},
				}
				engine.StartSubscriptionListener(nil)
				res, err := engine.GetFeatureStatusValue("")
				Expect(res).To(Equal(controlSuccessStatus))
				Ω(err).Should(BeNil())
			})

			It("Should return non nil err on failed call", func() {
				engine = &localpubsub.LocalPubSubEngine{
					LocalEngine: localEngineFailureMock,
					PubSub:      &PubSubConnMock{},
				}
				engine.StartSubscriptionListener(nil)
				_, err := engine.GetFeatureStatusValue("")
				Ω(err).ShouldNot(BeNil())
			})
			Context("go routine in StartSubscriptionListener", func() {
				var (
					pubsubCounter int                    = 0
					origReceiver  localpubsub.PSReceiver = nil
					wg            sync.WaitGroup
					wgGroupCount  int = 1
				)
				BeforeEach(func() {
					pubsubCounter = 0
					origReceiver = localpubsub.PubSubReceiver
					localpubsub.PubSubReceiver = func(s localpubsub.ReceiverInterface, toggleList map[string]*toggle.Feature) {
						pubsubCounter++
						if pubsubCounter <= wgGroupCount {
							wg.Done()
						}
					}
				})

				AfterEach(func() {
					pubsubCounter = 0
					localpubsub.PubSubReceiver = origReceiver
				})

				It("Should return non nil err on failed call", func() {
					wg.Add(wgGroupCount)
					engine = &localpubsub.LocalPubSubEngine{
						LocalEngine: localEngineFailureMock,
						PubSub:      &PubSubConnMock{},
					}
					engine.StartSubscriptionListener(nil)
					wg.Wait()
					Ω(pubsubCounter).Should(BeNumerically(">=", wgGroupCount))
				})
			})
		})
	})
})
