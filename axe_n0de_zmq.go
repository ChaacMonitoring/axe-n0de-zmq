package main

import (
  "fmt"
  "os"
  "flag"
  "strings"
  "strconv"
  "time"

  golzmq "github.com/abhishekkr/gol/golzmq"
  gollog "github.com/abhishekkr/gol/gollog"
  chaac_axe "github.com/ChaacMonitoring/axe-n0de-zmq/chaac_axe"
)

var (
  ip_address    = flag.String("ip", "0.0.0.0", "IP address of Chaac's Hand")
  req_port      = flag.Int("req-port", 9797, "what Socket PORT to run at")
  rep_port      = flag.Int("rep-port", 9898, "what Socket PORT to run at")
  put_interval  = flag.Int("put-interval", 10, "at what time interval(seconds) to push the results")
  get_interval  = flag.Int("get-interval", 15, "at what time interval(seconds) to pull the configs")
  host_id       = flag.String("host-id", "node", "prefix-ed namespace added to hostnames")

  hostname, _ = os.Hostname()
  host_signature = fmt.Sprintf("%s:%s", *host_id, hostname)
  request_socket = golzmq.ZmqRequestSocket(*ip_address, *req_port, *rep_port)
)


func get_config(){
  config_basekey := fmt.Sprintf("%s::config", host_signature)
  current_config, err := golzmq.ZmqRequest(request_socket, "read", config_basekey)
  if err != nil { gollog.Log_it(err.Error()) }
  if current_config == "" {
    config_set := []string{
      fmt.Sprintf("%s:put_interval,%d", config_basekey, *put_interval),
      fmt.Sprintf("%s:get_interval,%d", config_basekey, *get_interval),
    }
    current_config = strings.Join(config_set, "\n")
    golzmq.ZmqRequest(request_socket, "push", current_config)
  }
  for {
    config, _ := golzmq.ZmqRequest(request_socket, "read", config_basekey)
    fmt.Printf("WIP for reflecting %s", config)
    time.Sleep(time.Duration(*get_interval) * time.Second)
  }
}

func put_result(){
  var year, day, hour, min, sec int
  var month time.Month
  _str := strconv.Itoa
  for {
    year, month, day = time.Now().Date()
    hour, min, sec = time.Now().Clock()
    _result := chaac_axe.CheckResult()
    golzmq.ZmqRequest(request_socket, "push", "tsds",
      host_signature,
      _str(year), month.String(), _str(day),
      _str(hour), _str(min), _str(sec),
      _result)
    time.Sleep(time.Duration(*put_interval) * time.Second)
  }
}

func main(){
  flag.Parse()
  fmt.Printf("client ZeroMQ REP/REQ... at %d, %d", *req_port, *rep_port)

  fmt.Println("Checking out levigoTSDS based storage...")
  put_result()
}
