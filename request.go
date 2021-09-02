package main

import "time"

type Location struct {
	Address     string      `json:"address"`
	Coordinates Coordinates `json:"coordinates"`
}

type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Dimension struct {
	Depth  int64 `json:"depth"`
	Height int64 `json:"height"`
	Weight int64 `json:"weight"`
	Width  int64 `json:"width"`
}

type Package struct {
	Description string    `json:"description"`
	Dimensions  Dimension `json:"dimensions"`
	Name        string    `json:"name"`
	Price       int64     `json:"price"`
	Quantity    int64     `json:"quantity"`
}

type GrabExpressCreateQuotes struct {
	Destination Location  `json:"destination"`
	Origin      Location  `json:"origin"`
	Packages    []Package `json:"packages"`
	ServiceType string    `json:"serviceType"`
}

type Quote struct {
	Service struct {
		ID   int    `json:"id"`
		Type string `json:"type"`
		Name string `json:"name"`
	} `json:"service"`
	Currency struct {
		Code     string `json:"code"`
		Symbol   string `json:"symbol"`
		Exponent int    `json:"exponent"`
	} `json:"currency"`
	Amount            int `json:"amount"`
	EstimatedTimeline struct {
		Pickup  time.Time `json:"pickup"`
		Dropoff time.Time `json:"dropoff"`
	} `json:"estimatedTimeline"`
	Distance    int       `json:"distance"`
	Destination Location  `json:"destination"`
	Origin      Location  `json:"origin"`
	Packages    []Package `json:"packages"`
}

//////////////////////////////////////////////////////

type GrabExpressCreateDeliveries struct {
	MerchantOrderID string    `json:"merchantOrderID"`
	ServiceType     string    `json:"serviceType"`
	PaymentMethod   string    `json:"paymentMethod"`
	Packages        []Package `json:"packages"`
	Origin          Location  `json:"origin"`
	Destination     Location  `json:"destination"`
	Recipient       Recipient `json:"recipient"`
	Sender          Sender    `json:"sender"`
}

type Recipient struct {
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	SmsEnabled bool   `json:"smsEnabled"`
}

type Sender struct {
	FirstName   string `json:"firstName"`
	CompanyName string `json:"companyName"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	SmsEnabled  bool   `json:"smsEnabled"`
}

//////////////////////////////////////////////////////

type GrabExpressGetDeliveryByID struct {
	DeliveryID string `json:"deliveryID"`
}

type Timeline struct {
	Create    time.Time `json:"create"`
	Allocate  time.Time `json:"allocate"`
	Pickup    time.Time `json:"pickup"`
	Dropoff   time.Time `json:"dropoff"`
	Completed time.Time `json:"completed"`
}

type CashOnDelivery struct {
	Enable bool `json:"enable"`
	Amount int  `json:"amount"`
}

type AdvanceInfo struct {
	FailedReason string `json:"failedReason"`
}

//////////////////////////////////////////////////////

type GrabExpressCancelDelivery struct {
	DeliveryID string `json:"deliveryID"`
}

//////////////////////////////////////////////////////

type GrabExpressGenerateToken struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
	Scope        string `json:"scope"`
}

//////////////////////////////////////////////////////
