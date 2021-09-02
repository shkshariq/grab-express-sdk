package grab_express

import (
	"encoding/json"
	"fmt"
)

func (this GrabExpressCreateQuotes) APIName() string {
	return endpointQuotes
}

func (this GrabExpressCreateQuotes) Payload() (payload string) {

	b, err := json.Marshal(this)
	if err != nil {
		fmt.Println(err)
		return payload
	}

	return string(b)
}

func (this GrabExpressCreateQuotes) Params() map[string]string {
	var m = make(map[string]string)
	return m
}
