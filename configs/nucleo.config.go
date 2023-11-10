package configs

import (
	"time"

	"github.com/Bendomey/nucleo-go"
)

// NucleoConfig is the configuration for the nucleo service broker
var NucleoConfig = nucleo.Config{
	// Namespace of nodes to segment your nodes on the same network.
	Namespace: "nucleo-template",

	// Unique node identifier. Must be unique in a namespace. This is defaulted to the hostname of the machine once the broker is started.
	DiscoverNodeID: nil,

	// Default log level for built-in console logger. It can be overwritten in logger options above.
	// Available values: trace, debug, info, warn, error, fatal
	LogLevel: nucleo.LogLevelInfo,

	// Logger formatter defaults to text.
	// Available values: json, text.
	LogFormat: nucleo.LogFormatJSON,

	// Define transporter.
	// Available values: MEMORY, nats://localhost:4222, kafka://localhost:9092.
	// Note: These urls can be changed to point to your own nats/kafka cluster.
	Transporter: "MEMORY",

	// When using a custom transporter, you can define it here.
	// Here's an example: https://github.com/Bendomey/nucleo-go/blob/main/transit/nats/nats.go
	TransporterFactory: nil,

	// Number of seconds to send heartbeat packet to other nodes.
	HeartbeatFrequency: 5 * time.Second,

	// Number of seconds to wait before setting node to unavailable status.
	HeartbeatTimeout: 30 * time.Second,

	OfflineCheckFrequency: 20 * time.Second,

	OfflineTimeout: 10 * time.Minute,

	DontWaitForNeighbours: true,

	NeighboursCheckTimeout: 2 * time.Second,

	WaitForDependenciesTimeout: 2 * time.Second,

	// Limit of calling level. If it reaches the limit, broker will throw an MaxCallLevelError error. (Infinite loop protection)
	MaxCallLevel: 100,

	Metrics: false,
	// MetricsRate: 1,

	// Retry policy settings.
	RetryPolicy: nucleo.RetryPolicy{
		// Enable feature
		Enabled: false,

		// Count of retries
		Retries: 5,

		// First delay in milliseconds.
		Delay: 100,

		// Maximum delay in milliseconds.
		MaxDelay: 1000,

		// Backoff factor for delay. 2 means exponential backoff.
		Factor: 2,

		// A function to check failed requests.
		Check: func(err error) bool {
			return err != nil
		},
	},

	// Define serializer.
	// Available values: JSON.
	// Serializer: "JSON",

	// Available values: "RoundRobin", "Random"
	// Strategy: "RoundRobin",

	// Define registry strategy.
	// Define custom Strategy. By default, we use the random strategy.
	// Here's an example: https://github.com/Bendomey/nucleo-go/blob/main/transit/nats/nats.go
	StrategyFactory: nil,

	// Register custom middlewares
	Middlewares: []nucleo.Middlewares{},

	// Called after broker created.
	Created: func() {},

	// Called after broker started.
	Started: func() {},

	// Called after broker stopped.
	Stopped: func() {},
}
