package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Bendomey/nucleo-go/broker"
	"github.com/Bendomey/nucleo-template-project/configs"
)

func main() {
	nucleoBroker := broker.New(&configs.NucleoConfig)

	// list all services here
	nucleoBroker.PublishServices()

	nucleoBroker.Start()

	signalC := make(chan os.Signal, 1)
	signal.Notify(signalC, os.Interrupt, syscall.SIGTERM)

	<-signalC

	nucleoBroker.Stop()
}
