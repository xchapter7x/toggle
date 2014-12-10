package localpubsub

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
	"github.com/xchapter7x/toggle"
)

type ReceiverInterface interface {
	Receive() interface{}
}

type PSReceiver func(s ReceiverInterface, toggleList map[string]*toggle.Feature)

var PubSubReceiver PSReceiver = func(s ReceiverInterface, toggleList map[string]*toggle.Feature) {
	switch n := s.Receive().(type) {
	case redis.Message:
		toggleList[n.Channel].UpdateStatus(string(n.Data[:]))

	case error:
		fmt.Printf("error: %v\n", n)
		return
	}
}
