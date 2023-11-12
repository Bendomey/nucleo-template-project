package test

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/Bendomey/nucleo-go"
	"github.com/Bendomey/nucleo-go/broker"
	"github.com/Bendomey/nucleo-template-project/internal/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGreetingService(t *testing.T) {

	broker := broker.New(&nucleo.Config{LogLevel: "debug"})
	broker.PublishServices(services.GatewayApi, services.Greeter)

	broker.Start()

	defer broker.Stop()

	t.Run("Test '/api/greeter' endpoint", func(t *testing.T) {
		t.Run("should return with 'Hello World Nucleo'", func(t *testing.T) {
			response, err := http.Get("http://localhost:5001/api/greeter")

			require.NoError(t, err, "Invalid http request")
			want := "Hello World Nucleo"

			assert.Equal(t, bodyContent(response), want)
		})
	})
}

// bodyContent return the reponse body as string
func bodyContent(resp *http.Response) string {
	defer resp.Body.Close()
	bts, _ := ioutil.ReadAll(resp.Body)
	return string(bts)
}
