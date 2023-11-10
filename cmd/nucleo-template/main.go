package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Bendomey/nucleo-go/broker"
	"github.com/Bendomey/nucleo-template-project/configs"
	"github.com/Bendomey/nucleo-template-project/internal/services"
)

func main() {
	nucleoBroker := broker.New(&configs.NucleoConfig)

	// Register services
	nucleoBroker.PublishServices(
		services.Greeter,
	)

	nucleoBroker.Start()

	signalC := make(chan os.Signal, 1)
	signal.Notify(signalC, os.Interrupt, syscall.SIGTERM)

	<-signalC

	nucleoBroker.Stop()
}
