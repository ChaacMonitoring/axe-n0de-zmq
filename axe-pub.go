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

var (
  host        string
  channels    []string
  allChecks   map[string]string
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
func Publisher(checkKey string) {
  client := NewClient()
  hostname, _ := os.Hostname()
  for {
    checkCmd := allChecks[checkKey]
    message := fmt.Sprintf("%s : %s", checkKey, Executioner(checkCmd))
    channel := fmt.Sprintf("%s/%s", hostname, checkKey)
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
  allChecks = make(map[string]string)
  // Set up and parse command-line args.
  runtime.GOMAXPROCS(runtime.NumCPU())
  flag.StringVar(&host, "host", "127.0.0.1", "")
  //flag.StringVar(&channels, "channels", "", "")
  flag.Parse()
  // will be fetching checks,channels
  //channels = append(channels, "mem-free", "uptime")
  channels = append(channels, "uptime")

  for checkKey, checkCmd := range checks.BasicCheck {
    allChecks[checkKey] = checkCmd
  }
  for checkKey, checkCmd := range checks.MemoryCheck {
    channels = append(channels, checkKey)
    allChecks[checkKey] = checkCmd
  }

  // publisher
  RunPublishers()
}
