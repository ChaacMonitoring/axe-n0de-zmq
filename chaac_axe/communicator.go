package chaac_axe

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	gollog "github.com/abhishekkr/gol/gollog"
	golzmq "github.com/abhishekkr/gol/golzmq"
)

func GetConfig() {
	config_basekey := fmt.Sprintf("%s::config", hostSignature)
	current_config, err := golzmq.ZmqRequest(request_socket, "read", config_basekey)
	if err != nil {
		gollog.LogIt(LogFile, err.Error())
	}
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

func PutResult() {
	var year, day, hour, min, sec int
	var month time.Month
	_str := strconv.Itoa
	for {
		year, month, day = time.Now().Date()
		hour, min, sec = time.Now().Clock()
		result := CheckResult()
		golzmq.ZmqRequest(request_socket, "push", "tsds-csv",
			_str(year), _str(int(month)), _str(day),
			_str(hour), _str(min), _str(sec),
			result)
		fmt.Printf("Host: %s\nCSV: %s\n", hostSignature, result)
		time.Sleep(time.Duration(*put_interval) * time.Second)
	}
}
