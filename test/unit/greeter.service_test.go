package test

import (
	"testing"

	"github.com/Bendomey/nucleo-go"
	"github.com/Bendomey/nucleo-go/broker"
	"github.com/Bendomey/nucleo-template-project/internal/services"
	"github.com/stretchr/testify/assert"
)

func TestGreetingService(t *testing.T) {

	broker := broker.New(&nucleo.Config{LogLevel: "debug"})
	broker.PublishServices(services.Greeter)

	broker.Start()

	defer broker.Stop()

	t.Run("Test 'greeter.hello' action", func(t *testing.T) {
		t.Run("should return with 'Hello World Nucleo'", func(t *testing.T) {
			got := <-broker.Call("v1.greeter.hello", map[string]interface{}{})
			want := "Hello World Nucleo"

			assert.Equal(t, got.String(), want)
		})
	})
}
