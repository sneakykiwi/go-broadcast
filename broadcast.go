package broadcast

import "sync"

//Broadcaster is the root struct, an instance of this is used to register subscribers
type Broadcaster struct {
	cond        *sync.Cond
	subscribers map[interface{}]func(interface{})
	message     interface{}
	running     bool
}

//NewBroadcaster gives the broadcaster object to be used further in message
func NewBroadcaster() *Broadcaster {
	return &Broadcaster{
		cond:        sync.NewCond(&sync.RWMutex{}),
		subscribers: map[interface{}]func(interface{}){},
	}
}

//Subscribe takes in the id of the subscriber, which can be anything, and a function to run as the message callback
func (b *Broadcaster) Subscribe(id interface{}, f func(input interface{})) {
	b.subscribers[id] = f
}

//Unsubscribe removes a subscriber from the subscriber map
func (b *Broadcaster) Unsubscribe(id interface{}) {
	b.cond.L.Lock()
	delete(b.subscribers, id)
	b.cond.L.Unlock()
}

//Publish creates a new message and sends it to all subscribed functions
func (b *Broadcaster) Publish(message interface{}) {
	go func() {
		b.cond.L.Lock()

		b.message = message
		b.cond.Broadcast()
		b.cond.L.Unlock()
	}()
}

// Start the main broadcasting event
func (b *Broadcaster) Start() {
	b.running = true
	for b.running {
		b.cond.L.Lock()
		b.cond.Wait()
		go func() {
			for _, f := range b.subscribers {
				f(b.message) // publishes the message
			}
		}()
		b.cond.L.Unlock()
	}
}

// Stop broadcasting event
func (b *Broadcaster) Stop() {
	b.running = false
}
