package grab_express

func (this GrabExpressCancelDelivery) APIName() string {
	return endpointCancelBooking
}

func (this GrabExpressCancelDelivery) Payload() (payload string) {
	return
}

func (this GrabExpressCancelDelivery) Params() map[string]string {
	var m = make(map[string]string)
	m["deliveryID"] = this.DeliveryID
	return m
}
