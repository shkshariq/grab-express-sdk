package main

import (
	"encoding/json"
	"fmt"
)

func (this GrabExpressCreateDeliveries) APIName() string {
	return endpointPlaceBooking
}

func (this GrabExpressCreateDeliveries) Payload() (payload string) {

	b, err := json.Marshal(this)
	if err != nil {
		fmt.Println(err)
		return payload
	}

	return string(b)
}

func (this GrabExpressCreateDeliveries) Params() map[string]string {
	var m = make(map[string]string)
	return m
}
