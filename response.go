package main

type GrabExpressCreateQuotesResponse struct {
	Quotes      []Quote   `json:"quotes"`
	Origin      Location  `json:"origin"`
	Destination Location  `json:"destination"`
	Packages    []Package `json:"packages"`
}

type GrabExpressCreateDeliveriesResponse struct {
	DeliveryID      string      `json:"deliveryID"`
	MerchantOrderID string      `json:"merchantOrderID"`
	PaymentMethod   string      `json:"paymentMethod"`
	Quote           Quote       `json:"quote"`
	Sender          Sender      `json:"sender"`
	Recipient       Recipient   `json:"recipient"`
	Status          string      `json:"status"`
	TrackingURL     string      `json:"trackingURL"`
	Courier         interface{} `json:"courier"`
	Timeline        interface{} `json:"timeline"`
	Schedule        interface{} `json:"schedule"`
	CashOnDelivery  interface{} `json:"cashOnDelivery"`
	InvoiceNo       string      `json:"invoiceNo"`
	PickupPin       string      `json:"pickupPin"`
	AdvanceInfo     interface{} `json:"advanceInfo"`
}

type GrabExpressGetDeliveryByIDResponse struct {
	DeliveryID      string         `json:"deliveryID"`
	MerchantOrderID string         `json:"merchantOrderID"`
	PaymentMethod   string         `json:"paymentMethod"`
	Quote           Quote          `json:"quote"`
	Sender          Sender         `json:"sender"`
	Recipient       Recipient      `json:"recipient"`
	Status          string         `json:"status"`
	TrackingURL     string         `json:"trackingURL"`
	Courier         interface{}    `json:"courier"`
	Timeline        Timeline       `json:"timeline"`
	Schedule        interface{}    `json:"schedule"`
	CashOnDelivery  CashOnDelivery `json:"cashOnDelivery"`
	InvoiceNo       string         `json:"invoiceNo"`
	PickupPin       string         `json:"pickupPin"`
	AdvanceInfo     AdvanceInfo    `json:"advanceInfo"`
}

type GrabExpressCancelDeliveryResponse struct{}

type GrabExpressGenerateTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}
