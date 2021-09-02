package grab_express

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type Client struct {
	apiDomain  string
	authDomain string
	authToken  string
	Client     *http.Client
}

func NewClient(isProduction bool, authToken string) (client *Client) {
	client = &Client{}

	client.Client = http.DefaultClient
	client.authToken = authToken

	if isProduction {
		client.apiDomain = kProductionURL
		client.authDomain = kProductionAuthURL
	} else {
		client.apiDomain = kSandboxURL
		client.authDomain = kSandboxAuthURL
	}
	return client
}

func (this *Client) doRequest(method string, request GrabExpressParam, results interface{}, token bool) (err error) {

	var endpoint = request.APIName()
	params := request.Params()
	if params != nil {
		endpoint = strings.ReplaceAll(request.APIName(), "{$deliveryID}", params["deliveryID"])
	}
	var url string
	if token {
		url = this.authDomain + endpoint
	} else {
		url = this.apiDomain + endpoint
	}
	fmt.Println("REQUEST: ", method, url)
	var payload io.Reader
	if request.Payload() != "" {
		payload = strings.NewReader(request.Payload())
		fmt.Println("REQUEST PAYLOAD: ", payload)
	}

	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println("http.NewRequest.err", err)
		err = errors.New("Request failed")
		return err
	}

	if !token {
		req.Header.Set("Authorization", "Bearer "+this.authToken)
	}

	req.Header.Set("Content-Type", "no-cache")
	req.Header.Set("cache-control", kContentType)

	resp, err := this.Client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		fmt.Println("httpClient.Do.err", err)
		err = errors.New("Request failed")
		return err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ioutil.ReadAll.err", err)
		err = errors.New("Request failed")
		return err
	}

	fmt.Println("RESPONSE.StatusCode:", resp.StatusCode)
	if resp.StatusCode != 200 {
		fmt.Println("RESPONSE.Data:", string(data))
		err = errors.New("Request failed")
		return err
	}
	fmt.Println("RESPONSE.Data:", string(data))

	err = json.Unmarshal(data, results)
	if err != nil {
		fmt.Println("json.Unmarshal.err", err)
		err = errors.New("Request failed")
		return err
	}

	return err
}

//https://developer-beta.stg-myteksi.com/docs/grab-express/#create-quotes
func (this *Client) CreateQuotes(request GrabExpressCreateQuotes) (response *GrabExpressCreateQuotesResponse, err error) {
	err = this.doRequest("POST", request, &response, false)
	return response, err
}

//https://developer-beta.stg-myteksi.com/docs/grab-express/#create-deliveries
func (this *Client) CreateDeliveries(request GrabExpressCreateDeliveries) (response *GrabExpressCreateDeliveriesResponse, err error) {
	err = this.doRequest("POST", request, &response, false)
	return response, err
}

//https://developer-beta.stg-myteksi.com/docs/grab-express/#get-delivery-by-id
func (this *Client) GetDeliveryByID(request GrabExpressGetDeliveryByID) (response *GrabExpressGetDeliveryByIDResponse, err error) {
	err = this.doRequest("GET", request, &response, false)
	return response, err
}

//https://developer-beta.stg-myteksi.com/docs/grab-express/#cancel-delivery
func (this *Client) CancelDelivery(request GrabExpressCancelDelivery) (response *GrabExpressCancelDeliveryResponse, err error) {
	err = this.doRequest("DELETE", request, &response, false)
	return response, err
}

//https://api.stg-myteksi.com/grabid/v1/oauth2/token
func (this *Client) GenerateToken(request GrabExpressGenerateToken) (response *GrabExpressGenerateTokenResponse, err error) {
	err = this.doRequest("POST", request, &response, true)
	return response, err
}
