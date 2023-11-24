package services

import "github.com/Bendomey/nucleo-go"

var Greeter = nucleo.ServiceSchema{
	Name: "greeter",

	// All your settings goes here
	Settings: map[string]interface{}{},

	// Dependencies are the services that this service depends on
	Dependencies: []string{},

	Version: "v1",

	// Actions are the actions that this service will perform
	Actions: []nucleo.Action{
		// action to say hello
		{
			// Name of action
			Name: "hello",

			// Description of action. This really has no use in internal services. It just helps with code readability.
			Description: "Say hello to the world",

			// validations for your params should be here.
			Params: map[string]interface{}{},

			// Handler is the function that will be called when this action is called
			Handler: func(context nucleo.Context, params nucleo.Payload) interface{} {
				return "Hello World Nucleo"
			},
		},

		{
			// Name of action
			Name: "say",

			// Description of action. This really has no use in internal services. It just helps with code readability.
			Description: "Say hello to the world",

			// validations for your params should be here.
			Params: map[string]interface{}{
				"something": "required",
			},

			// Handler is the function that will be called when this action is called
			Handler: func(context nucleo.Context, params nucleo.Payload) interface{} {
				return "Say " + params.Get("something").String()
			},
		},
	},
}
