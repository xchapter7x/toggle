package localpubsub

import (
	"fmt"
	"os"

	"github.com/garyburd/redigo/redis"
	"github.com/xchapter7x/toggle/engines/localengine"
	"github.com/xchapter7x/toggle/engines/storageinterface"
)

func NewLocalPubSubEngine(pubsub pubsubInterface) storageinterface.StorageEngine {
	le := &localengine.LocalEngine{
		Getenv: os.Getenv,
	}
	engine := &LocalPubSubEngine{
		LocalEngine: le,
		PubSub:      pubsub,
	}
	engine.StartSubscriptionListener()
	return engine
}

type LocalPubSubEngine struct {
	LocalEngine *localengine.LocalEngine
	PubSub      pubsubInterface
	quit        chan bool
}

func (s *LocalPubSubEngine) Close() (err error) {
	s.quit <- true
	return
}

func pubsubReciever(s *LocalPubSubEngine) {
	for {
		switch n := s.PubSub.Receive().(type) {
		case redis.Message:
			fmt.Printf("Message: %s %s\n", n.Channel, n.Data)
		case redis.Subscription:
			fmt.Printf("Subscription: %s %s %d\n", n.Kind, n.Channel, n.Count)
			if n.Count == 0 {
				return
			}
		case error:
			fmt.Printf("error: %v\n", n)
			return
		}
	}
}

func (s *LocalPubSubEngine) StartSubscriptionListener() {
	if s.quit == nil {
		s.quit = make(chan bool)

		go func() {
			for {
				select {
				case <-s.quit:
					return
				default:
					pubsubReciever(s)
				}
			}
		}()
	}
}

func (s *LocalPubSubEngine) GetFeatureStatusValue(featureSignature string) (status string, err error) {
	s.PubSub.Subscribe(featureSignature)
	status, err = s.LocalEngine.GetFeatureStatusValue(featureSignature)
	return
}
