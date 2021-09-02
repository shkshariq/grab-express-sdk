package main

import (
	ge "gitlab.com/dishserve/go/sdk/grab-express/grab-express"
	"testing"
)

//https://developer-beta.stg-myteksi.com/docs/grab-express/#authentication-method
var token = `eyJhbGciOSUzI1NiIsImtpZCI6Il9kZWZhdWx0IiwidHlwIjoiSldUIn0.eyJhdWQiOiIwZmRiMjZhMzExMDk0MDJmODE1NDI2MGNlYmJkM2VmMiIsImV4cCI6MTYxNTczNjY3NSwiaWF0IjoxNjE1NjUwMjc1LCJpc3MiOiJodHRwczovL2lkcC5ncmFiLmNvbSIsImp0aSI6ImdjUnppQUdsUmN5bnlPTUNROHlwd2ciLCJuYmYiOjE2MTU2NTAwOTUsInBpZCI6IjRjZTljZDhiLTVhMjktNDIxNS04MjBmLTQ2YWMyODJkMTM2YyIsInBzdCI6MSwic2NwIjoiW1wiN2MxNDk3NGQzZDBlNDYyYjgzZDM1OGYwNTViZjdiYzZcIl0iLCJzdWIiOiJUV09fTEVHR0VEX09BVVRIIiwic3ZjIjoiIiwidGtfdHlwZSI6ImFjY2VzcyJ9.nqw4Ttytx136OFk3mkCI7783A5gxKocbBSdTzMrkiPgBVmVwE3GzyrglthJ5pJ1WmKYK_hGxgrDm6TCefRKvcDk0X6zjnH6VRUiqrqKH8GwZXnL7yoX3YS4AG_z31MWdy6uLMau71TkGFmO0xkD4sAuXwV92NmY5smWssLSGyfu-TAjVqYy4XESDvwCjlVS4KLIOJammzXGXi1Ipa6jhBgdBNn_qPUnZYjhhGxG-lOqlEOQqh18zmdRz2MlVKw-QmwovhuXtMD-u_t7u04uJ0KZafUn00jMMwk7J8KzJFDQEBzW90UnqhxcD25pou8q9SFZA9VgKO0XBGTofTqY7uA`

func TestGetQuote(t *testing.T) {

	client := ge.NewClient(false, token)

	// Preparing Request
	req := ge.GrabExpressCreateQuotes{}
	req.ServiceType = "INSTANT"
	req.Packages = []ge.Package{
		ge.Package{
			Name:        "Fish Burger",
			Description: "Fish Burger with mayonnaise sauce",
			Quantity:    2,
			Price:       5,
			Dimensions: ge.Dimension{
				Height: 0,
				Width:  0,
				Depth:  0,
				Weight: 0,
			},
		},
		ge.Package{
			Name:        "Truffle Fries",
			Description: "Thin cut deep-fried potatoes topped with truffle oil",
			Quantity:    2,
			Price:       4,
			Dimensions: ge.Dimension{
				Height: 0,
				Width:  0,
				Depth:  0,
				Weight: 0,
			},
		},
	}
	req.Origin = ge.Location{
		Address: "1 IJK View, Singapore 018936",
		Coordinates: ge.Coordinates{
			Latitude:  1.2345678,
			Longitude: 3.4567890,
		},
	}
	req.Destination = ge.Location{
		Address: "1 ABC St, Singapore 078881",
		Coordinates: ge.Coordinates{
			Latitude:  1.2345678,
			Longitude: 3.4567890,
		},
	}

	res, err := client.CreateQuotes(req)
	if err != nil {
		t.Errorf("Test failed: %v", err)
		return
	}
	t.Logf("Test passed: %v", res)
}

func TestCreateDeliveries(t *testing.T) {

	client := ge.NewClient(false, token)

	// Preparing Request
	req := ge.GrabExpressCreateDeliveries{}
	req.MerchantOrderID = "1ac1fa2f-880a-43d3-8476-7cc6e99e40f6"
	req.ServiceType = "INSTANT"
	req.PaymentMethod = "CASHLESS"
	req.Origin = ge.Location{
		Address: "1 IJK View, Singapore 018936",
		Coordinates: ge.Coordinates{
			Latitude:  1.2345678,
			Longitude: 3.4567890,
		},
	}
	req.Destination = ge.Location{
		Address: "1 ABC St, Singapore 078881",
		Coordinates: ge.Coordinates{
			Latitude:  1.2345678,
			Longitude: 3.4567890,
		},
	}
	req.Recipient = ge.Recipient{
		FirstName:  "John",
		LastName:   "Tan",
		Email:      "john@gmail.com",
		Phone:      "91526655",
		SmsEnabled: true,
	}
	req.Sender = ge.Sender{
		FirstName:   "Jewel Changi Branch",
		CompanyName: "Shake Shack",
		Email:       "ssburger@gmail.com",
		Phone:       "0124355834",
		SmsEnabled:  true,
	}

	req.Packages = []ge.Package{
		ge.Package{
			Name:        "Fish Burger",
			Description: "Fish Burger with mayonnaise sauce",
			Quantity:    2,
			Price:       5,
			Dimensions: ge.Dimension{
				Height: 0,
				Width:  0,
				Depth:  0,
				Weight: 0,
			},
		},
		ge.Package{
			Name:        "Truffle Fries",
			Description: "Thin cut deep-fried potatoes topped with truffle oil",
			Quantity:    2,
			Price:       4,
			Dimensions: ge.Dimension{
				Height: 0,
				Width:  0,
				Depth:  0,
				Weight: 0,
			},
		},
	}

	res, err := client.CreateDeliveries(req)
	if err != nil {
		t.Errorf("Test failed: %v", err)
		return
	}
	t.Logf("Test passed: %v", res)
}

func TestGetDeliveryByID(t *testing.T) {

	client := ge.NewClient(false, token)

	// Preparing Request
	req := ge.GrabExpressGetDeliveryByID{}
	req.DeliveryID = "IN-2-00D2B1CBMWBO7JND48J5"

	res, err := client.GetDeliveryByID(req)
	if err != nil {
		t.Errorf("Test failed: %v", err)
		return
	}
	t.Logf("Test passed: %v", res)
}

func TestCancelDelivery(t *testing.T) {

	client := ge.NewClient(false, token)

	// Preparing Request
	req := ge.GrabExpressCancelDelivery{}
	req.DeliveryID = "IN-2-00D2B1CBMWBO7JND48J5"

	res, err := client.CancelDelivery(req)
	if err != nil {
		t.Errorf("Test failed: %v", err)
		return
	}
	t.Logf("Test passed: %v", res)
}
