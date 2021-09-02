package main

import (
	"encoding/json"
	"fmt"
)

func (this GrabExpressGenerateToken) APIName() string {
	return endpointTokenInformation
}

func (this GrabExpressGenerateToken) Payload() (payload string) {
	b, err := json.Marshal(this)
	if err != nil {
		fmt.Println(err)
		return payload
	}

	return string(b)
}

func (this GrabExpressGenerateToken) Params() map[string]string {
	var m = make(map[string]string)
	return m
}
