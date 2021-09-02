package main

const kContentType = "application/json"

// Auth Server Creds
const (
	kProductionURL     = ""
	kProductionAuthURL = ""
)

// GrabExpress Server Creds
const (
	kSandboxURL     = "https://partner-api.stg-myteksi.com/grab-express-sandbox"
	kSandboxAuthURL = "https://api.stg-myteksi.com/grabid"
)

const (
	endpointTokenInformation = "/v1/oauth2/token"
)

const (
	endpointQuotes           = "/v1/deliveries/quotes"
	endpointPlaceBooking     = "/v1/deliveries"
	endpointCancelBooking    = "/v1/deliveries/{$deliveryID}"
	endpointOrderInformation = "/v1/deliveries/{$deliveryID}"
)

type GrabExpressParam interface {
	// grab_express endpoint
	APIName() string

	// grab_express json payload
	Payload() string

	// grab_express params
	Params() map[string]string
}
