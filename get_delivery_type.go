package main

func (this GrabExpressGetDeliveryByID) APIName() string {
	return endpointOrderInformation
}

func (this GrabExpressGetDeliveryByID) Payload() (payload string) {
	return
}

func (this GrabExpressGetDeliveryByID) Params() map[string]string {
	var m = make(map[string]string)
	m["deliveryID"] = this.DeliveryID
	return m
}
