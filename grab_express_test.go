package grab_express

import (
	"testing"
)

//https://developer-beta.stg-myteksi.com/docs/grab-express/#authentication-method
var token = `eyJhbGciiJSUzI1NiIsImtpZCI6Il9kZWZhdWx0IiwidHlwIjoiSldUIn0.eyJhdWQiOiIwZmRiMjZhMzExMDk0MDJmODE1NDI2MGNlYmJkM2VmMiIsImV4cCI6MTYzMDY3MzcwNCwiaWF0IjoxNjMwNTg3MzA0LCJpc3MiOiJodHRwczovL2lkcC5ncmFiLmNvbSIsImp0aSI6InphVUdaUk5rVDMyYU5jNm5LVjBGd3ciLCJuYmYiOjE2MzA1ODcxMjQsInBpZCI6IjRjZTljZDhiLTVhMjktNDIxNS04MjBmLTQ2YWMyODJkMTM2YyIsInBzdCI6MSwic2NwIjoiW1wiN2MxNDk3NGQzZDBlNDYyYjgzZDM1OGYwNTViZjdiYzZcIl0iLCJzdWIiOiJUV09fTEVHR0VEX09BVVRIIiwic3ZjIjoiIiwidGtfdHlwZSI6ImFjY2VzcyJ9.q5FIWJ1mj5i2YhujDvdm_KkwAmiI0lL-Mx9HOFHvmFAbCi2wpb1tbZs0AG9E7IAo0gI6WsFNcx1wCNdI6DtZ0eCWkLVHxRvptbMmWZLV8x-pZuOvsa1cCSQlqOFKO_Qq5-ADZDMQYKlBm6yl3_iYECeAmRxDDH5gqdX3O2Oxun_w97RhJk-xTu0D58Bqspcz3B7D9HLEyTAiRF9BR8CDP0xwCthDThd-kFWluAYfzjTj5mox6-rn5PiVidPSBCx4e1iQngU91lj410lu4wceZSGX5UhcSpy0azp0_JRMs-5ZFvcNT2463QeYFHhDnxxsityYXc_YuRgHaCXuLSDZVQ`

func TestGetQuote(t *testing.T) {

	client := NewClient(false, token)

	// Preparing Request
	req := GrabExpressCreateQuotes{}
	req.ServiceType = "INSTANT"
	req.Packages = []Package{
		Package{
			Name:        "Fish Burger",
			Description: "Fish Burger with mayonnaise sauce",
			Quantity:    2,
			Price:       5,
			Dimensions: Dimension{
				Height: 0,
				Width:  0,
				Depth:  0,
				Weight: 0,
			},
		},
		Package{
			Name:        "Truffle Fries",
			Description: "Thin cut deep-fried potatoes topped with truffle oil",
			Quantity:    2,
			Price:       4,
			Dimensions: Dimension{
				Height: 0,
				Width:  0,
				Depth:  0,
				Weight: 0,
			},
		},
	}
	req.Origin = Location{
		Address: "1 IJK View, Singapore 018936",
		Coordinates: Coordinates{
			Latitude:  1.2345678,
			Longitude: 3.4567890,
		},
	}
	req.Destination = Location{
		Address: "1 ABC St, Singapore 078881",
		Coordinates: Coordinates{
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

	client := NewClient(false, token)

	// Preparing Request
	req := GrabExpressCreateDeliveries{}
	req.MerchantOrderID = "1ac1fa2f-880a-43d3-8476-7cc6e99e40f6"
	req.ServiceType = "INSTANT"
	req.PaymentMethod = "CASHLESS"
	req.Origin = Location{
		Address: "1 IJK View, Singapore 018936",
		Coordinates: Coordinates{
			Latitude:  1.2345678,
			Longitude: 3.4567890,
		},
	}
	req.Destination = Location{
		Address: "1 ABC St, Singapore 078881",
		Coordinates: Coordinates{
			Latitude:  1.2345678,
			Longitude: 3.4567890,
		},
	}
	req.Recipient = Recipient{
		FirstName:  "John",
		LastName:   "Tan",
		Email:      "john@gmail.com",
		Phone:      "91526655",
		SmsEnabled: true,
	}
	req.Sender = Sender{
		FirstName:   "Jewel Changi Branch",
		CompanyName: "Shake Shack",
		Email:       "ssburger@gmail.com",
		Phone:       "0124355834",
		SmsEnabled:  true,
	}

	req.Packages = []Package{
		Package{
			Name:        "Fish Burger",
			Description: "Fish Burger with mayonnaise sauce",
			Quantity:    2,
			Price:       5,
			Dimensions: Dimension{
				Height: 0,
				Width:  0,
				Depth:  0,
				Weight: 0,
			},
		},
		Package{
			Name:        "Truffle Fries",
			Description: "Thin cut deep-fried potatoes topped with truffle oil",
			Quantity:    2,
			Price:       4,
			Dimensions: Dimension{
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

	client := NewClient(false, token)

	// Preparing Request
	req := GrabExpressGetDeliveryByID{}
	req.DeliveryID = "IN-2-00D2B1CBMWBO7JND48J5"

	res, err := client.GetDeliveryByID(req)
	if err != nil {
		t.Errorf("Test failed: %v", err)
		return
	}
	t.Logf("Test passed: %v", res)
}

func TestCancelDelivery(t *testing.T) {

	client := NewClient(false, token)

	// Preparing Request
	req := GrabExpressCancelDelivery{}
	req.DeliveryID = "IN-2-00D2B1CBMWBO7JND48J5"

	res, err := client.CancelDelivery(req)
	if err != nil {
		t.Errorf("Test failed: %v", err)
		return
	}
	t.Logf("Test passed: %v", res)
}
