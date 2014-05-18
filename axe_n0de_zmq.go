package main

import (
	"fmt"

	gollog "github.com/abhishekkr/gol/gollog"

	chaac_axe "github.com/ChaacMonitoring/axe-n0de-zmq/chaac_axe"
)

func main() {
	defer gollog.CloseLogFile(chaac_axe.LogFile)
	chaac_axe.PutResult()
	fmt.Println("Pushed Check Status.")
}
