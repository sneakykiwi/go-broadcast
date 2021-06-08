package broadcast_test

import (
	"fmt"
	"github.com/sneakykiwi/go-broadcast"
	"testing"
	"time"
)

var brdcstr = broadcast.NewBroadcaster()

func TestBroadcaster(t *testing.T){
	messageToaster := func(message interface{}) {
		fmt.Printf("[New Message]: %v\n", message)
	}
	unwillingReceiver := func(message interface{}) {
		fmt.Println("Do not disturb!")
	}
	brdcstr.Subscribe(1, messageToaster)
	brdcstr.Subscribe(2, messageToaster)
	brdcstr.Subscribe(3, unwillingReceiver)

	go brdcstr.Start()

	brdcstr.Publish("Hello!")

	time.Sleep(time.Second)
	brdcstr.Unsubscribe(3)
	brdcstr.Publish("Goodbye!")
}