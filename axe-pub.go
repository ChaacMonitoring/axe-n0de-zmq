package main

import (
  // go packages
	"flag"
	"runtime"
  "time"
  "fmt"
  "os"
  "os/exec"
  "log"
  // in-house packages
  "./zmq_pub_sub"
  "./checks"
)

var all_checks =  make(map[string]string)

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

// runner
func Executioner(Cmd string) (s string) {
  out, err := exec.Command(Cmd).Output()
  if err != nil {
    println("ERROR:")
    log.Fatal(err)
  }
  return string(out)
}

// Publisher forever
func Publisher(check_key string) {
	client := NewClient()
  hostname, _ := os.Hostname()
	for {
    check_cmd := all_checks[check_key]
    message := fmt.Sprintf("%s : %s", check_key, Executioner(check_cmd))
    channel := fmt.Sprintf("%s/%s", hostname, check_key)
		client.Publish(channel, message)
    time.Sleep(60 * time.Second)
	}
}

// creates goroutines, for each Check
func RunPublishers() {
	for _, channel := range channels {
		//go Publisher(channel) // need to goroutine
    println("publishing...", channel)
		Publisher(channel)
	}
}

func main() {
	// Set up and parse command-line args.
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.StringVar(&host, "host", "127.0.0.1", "")
	//flag.StringVar(&channels, "channels", "", "")
	flag.Parse()
  // will be fetching checks,channels
  //channels = append(channels, "mem-free", "uptime")
  channels = append(channels, "uptime")

  for check_key, check_cmd := range checks.BasicCheck {
    all_checks[check_key] = check_cmd
  }
  for check_key, check_cmd := range checks.MemoryCheck {
    channels = append(channels, check_key)
    all_checks[check_key] = check_cmd
  }

	// publisher
  RunPublishers()
}
