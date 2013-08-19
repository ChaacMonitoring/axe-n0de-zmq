package main

import (
  // go packages
	"flag"
  "fmt"
  "time"
	"runtime"
	"strconv"
  "os"
  // in-house packages
  "./zmq_pub_sub"
  "./checks"
)

var (
	host        string
	channels    []string
)

// new pubsub client instance
func NewClient() pubsub.Client {
	var client pubsub.Client
	client = pubsub.NewZMQClient(host)
	return client
}

// Subsribe forever
func Subscriber(channel string) {
	client := NewClient()
  client.Subscribe(channel)
	last := time.Now()
	messages := 0
  fmt.Println(last)
	for {
		data := client.Receive()
    fmt.Println(data)
		messages += 1
		now := time.Now()
		if now.Sub(last).Seconds() > 1 {
			fmt.Println(messages, "msg/sec")
			client.Publish("metrics", strconv.Itoa(messages))
			last = now
			messages = 0
		}
	}
}

// creates goroutines, for each Check
func RunSubscribers() {
	for _, channel := range channels {
		// go Subscriber(channel)
    println("subscribing...", channel)
		Subscriber(channel)
	}
}

func main() {
	// Set up and parse command-line args.
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.StringVar(&host, "host", "127.0.0.1", "")
	flag.Parse()
  // will be fetching checks,channels

  hostname, _ := os.Hostname()
  for check_key, _ := range checks.BasicCheck {
    channel_name := fmt.Sprintf("%s/%s", hostname, check_key)
    channels = append(channels, channel_name)
  }
  for check_key, _ := range checks.MemoryCheck {
    channel_name := fmt.Sprintf("%s/%s", hostname, check_key)
    channels = append(channels, channel_name)
  }

	// publisher
  RunSubscribers()
}
