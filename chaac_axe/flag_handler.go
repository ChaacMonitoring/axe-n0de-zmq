package chaac_axe

import (
	"flag"
	"fmt"
	"os"

	gollog "github.com/abhishekkr/gol/gollog"
	golzmq "github.com/abhishekkr/gol/golzmq"
)

var (
	ip_address   = flag.String("ip", "0.0.0.0", "IP address of Chaac's Hand")
	req_port     = flag.Int("req-port", 9797, "what Socket PORT to run at")
	rep_port     = flag.Int("rep-port", 9898, "what Socket PORT to run at")
	put_interval = flag.Int("put-interval", 10, "at what time interval(seconds) to push the results")
	get_interval = flag.Int("get-interval", 15, "at what time interval(seconds) to pull the configs")
	host_id      = flag.String("host-id", "node", "prefix-ed namespace added to hostnames")
	logFilename  = flag.String("logfile", "", "path to log to")

	hostname, _    = os.Hostname()
	hostSignature  = fmt.Sprintf("%s:%s", *host_id, hostname)
	request_socket = golzmq.ZmqRequestSocket(*ip_address, *req_port, *rep_port)

	LogFile *os.File
)

func init() {
	flag.Parse()
	fmt.Printf("client ZeroMQ REP/REQ... at %d, %d\n", *req_port, *rep_port)

	if *logFilename == "" {
		*logFilename = fmt.Sprintf("%s.log", os.Args[0])
	}
	LogFile = gollog.OpenLogFile(*logFilename)
}
