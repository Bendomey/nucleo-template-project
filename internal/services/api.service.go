package services

import (
	"fmt"

	"github.com/Bendomey/awesome-nucleo/gateway"
	"github.com/Bendomey/nucleo-go"
	"github.com/Bendomey/nucleo-go/errors"
	"github.com/Bendomey/nucleo-go/payload"
	"github.com/Bendomey/nucleo-go/serializer"
	"github.com/gin-gonic/gin"
)

var GatewayMixin = gateway.NewGatewayMixin(gateway.GatewayMixin{
	Authenticate: &authenticateHandler,
	Authorize:    &authorizeHandler,
})

var GatewayApi = nucleo.ServiceSchema{
	Name: "api",
	Mixins: []nucleo.Mixin{
		GatewayMixin,
	},
	Settings: map[string]interface{}{

		// Exposed Port
		"port": 5001,

		// Exposed IP
		"ip": "0.0.0.0",

		// base path
		"path": "/",

		// Global middlewares. Applied to all routes.
		"use": []gin.HandlerFunc{},

		// Routes
		"routes": []gateway.Route{
			{
				Name: "basic",

				Whitelist: []string{"**"},

				Path: "/api",

				// Route-level middlewares
				Use: []gin.HandlerFunc{},

				// Mapping policy setting
				// Available values: "all", "restrict"
				MappingPolicy: gateway.MappingPolicyRestrict,

				Aliases: map[string]string{
					"GET /greeter":      "v1.greeter.hello",
					"POST /greeter/say": "v1.greeter.say",
				},

				// Before call hook. You can check the request.
				OnBeforeCall: &OnBeforeCallHandler,

				//  After call hook. You can modify the data.
				OnAfterCall: &OnAfterCallHandler,

				// Enable authentication. Implement the logic into `authenticate` method.
				Authorization: false,

				// Enable authentication. Implement the logic into `authenticate` method.
				Authentication: false,
			},
		},

		// Log each request (default to "info" level)
		"logRequest": nucleo.LogLevelDebug,

		// Log the request ctx.params (default to "debug" level)
		"logRequestParams": nucleo.LogLevelInfo,

		// Log each response (default to "info" level)
		"logResponse": nucleo.LogLevelInfo,

		// Log the response data (default to disable)
		"logResponseData": nucleo.LogLevelInfo,

		// If set to true, it will log 4xx client errors, as well
		"log4XXResponses": nucleo.LogLevelInfo,

		// Log the route registration/aliases related activity
		"logRouteRegistration": nucleo.LogLevelInfo,

		"onError": func(context *gin.Context, response nucleo.Payload) {
			jsonSerializer := serializer.JSONSerializer{}

			nucleoError := response.Value().(errors.NucleoError)

			responsePayload := payload.New(map[string]interface{}{
				"message": nucleoError.Message,
				"type":    nucleoError.Type,
				"code":    nucleoError.Code,
				"data":    nucleoError.Data,
			})
			json := jsonSerializer.PayloadToBytes(responsePayload)

			context.Writer.WriteHeader(nucleoError.Code)
			context.Writer.Write(json)
		},
	},
}

/**
 * Authenticate the request. It check the `Authorization` token value in the request header.
 * Check the token value & resolve the user by the token.
 * The resolved user will be available in `context.Meta().Get("user")`
 *
 * PLEASE NOTE, IT'S JUST AN EXAMPLE IMPLEMENTATION. DO NOT USE IN PRODUCTION!
 */
var authenticateHandler = func(context nucleo.Context, ginContext *gin.Context, alias string) interface{} {
	fmt.Println("authenticate called")

	return map[string]interface{}{
		"name":  "Benjamin",
		"token": "ds.kljvbajh akjsbfkhas fjnas",
	}
}

/**
 * Authorize the request. Check that the authenticated user has right to access the resource.
 *
 * PLEASE NOTE, IT'S JUST AN EXAMPLE IMPLEMENTATION. DO NOT USE IN PRODUCTION!
 */
var authorizeHandler = func(context nucleo.Context, ginContext *gin.Context, alias string) {

	// Get the authenticated user.
	authenticatedUser := context.Meta().Get("user")

	fmt.Println("authorize called. user: ", authenticatedUser)
}

var OnBeforeCallHandler = func(context nucleo.Context, ginContext *gin.Context, route gateway.Route, alias string) {
	fmt.Println("on before call handler")
}

var OnAfterCallHandler = func(context nucleo.Context, ginContext *gin.Context, route gateway.Route, response nucleo.Payload) {
	fmt.Println("on after call handler")
}
