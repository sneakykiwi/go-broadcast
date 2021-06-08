# Broadcaster

This module gives you the ability to publish messages in a shared and thread-safe manner, it is designed to work similarly
to how `Redis Pub/Sub` works. 

The module has 100% versatility as it is able to accept all types of callbacks, you could use this for publishing any sorts
of events to all subscribed goroutines. 

**<h3>No channels used</h3>**

# Examples

```go
func main(){
    messageToaster := func(message interface{}) {
        fmt.Printf("[New Message]: %v\n", message)
    }
    unwillingReceiver := func(message interface{}) {
        fmt.Println("Do not disturb!")
    }
    broadcaster := SetupBroadcaster()
    broadcaster.Subscribe(1, messageToaster)
    broadcaster.Subscribe(2, messageToaster)
    broadcaster.Subscribe(3, unwillingReceiver)

    go broadcaster.Start()

    broadcaster.Publish("Hello!")

    time.Sleep(time.Second)
    broadcaster.Unsubscribe(3)
    broadcaster.Publish("Goodbye!")
    // [New Message]: Hello!
    // Do not disturb!
    // [New Message]: Hello!
    // [New Message]: Goodbye!
    // [New Message]: Goodbye!
}
```


# Benchmarks
```json
pkg: github.com/sneakykiwi/go-broadcast
cpu: Intel(R) Core(TM) i5-9400F CPU @ 2.90GHz
BenchmarkBroadcaster100-6       1000000000
BenchmarkBroadcaster500-6       1000000000
BenchmarkBroadcaster1000-6      1000000000
PASS
ok      github.com/sneakykiwi/go-broadcast      1.289s
```