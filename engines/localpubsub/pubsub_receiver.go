package localpubsub

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func PubSubReceiver(s pubsubInterface) {
	switch n := s.Receive().(type) {
	case redis.Message:
		fmt.Printf("Message: %s %s\n", n.Channel, n.Data)

	case error:
		fmt.Printf("error: %v\n", n)
		return
	}
}
