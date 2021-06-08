package broadcast_test

import (
	"testing"
)

func benchmarkBroadcaster(j int, b *testing.B){
	messageToaster := func(message interface{}) {
		//fmt.Printf("[New Message]: %v\n", message)
	}
	unwillingReceiver := func(message interface{}) {
		//fmt.Println("Do not disturb!")
	}
	brdcstr.Subscribe(1, messageToaster)
	brdcstr.Subscribe(2, messageToaster)
	brdcstr.Subscribe(3, unwillingReceiver)

	go brdcstr.Start()


	go func() {
		for i := 0; i < j; i++ {
			brdcstr.Publish("Hello!")
		}
	}()
}

func BenchmarkBroadcaster100(b *testing.B) { benchmarkBroadcaster(100, b) }
func BenchmarkBroadcaster500(b *testing.B) { benchmarkBroadcaster(500, b) }
func BenchmarkBroadcaster1000(b *testing.B) { benchmarkBroadcaster(1000, b) }