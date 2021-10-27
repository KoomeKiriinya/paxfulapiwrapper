package paxful

import (
	"errors"
	"os"
)

const (
	BaseURLv1 = "https://api.paxful.com/paxful/v1"
)

type Client struct {
	BaseURL string
}

func NewClient() *Client {
	return &Client{
		BaseURL: BaseURLv1,
	}
}

func (c *Client) Auth() (string, error) {
	auth_details := Auth{
		Client_id:     os.Getenv("CLIENT_ID"),
		Client_secret: os.Getenv("CLIENT_SECRET"),
		Grant_type:    os.Getenv("GRANT_TYPE"),
	}

	auth_url := os.Getenv("AUTH_URL")
	auth_body := "client_id=" + auth_details.Client_id + "&client_secret=" + auth_details.Client_secret + "&grant_type=" + auth_details.Grant_type
	res_body, err := post(auth_url, "", auth_body)
	if err != nil {
		return "", errors.New("error occured : " + err.Error())
	}

	return res_body["access_token"].(string), nil
}

// Currency

func (c *Client) CurrencyRates() (map[string]interface{}, error) {
	req_url := c.BaseURL + "/currency/rates"
	token, err := c.Auth()
	if err != nil {
		return map[string]interface{}{}, errors.New("error occured : " + err.Error())
	}
	response, err := post(req_url, token, "")
	if err != nil {
		return map[string]interface{}{}, errors.New("error occured : " + err.Error())
	}
	return response, nil
}

func (c *Client) CurrencyList() (map[string]interface{}, error) {
	req_url := c.BaseURL + "/currency/list"
	token, err := c.Auth()
	if err != nil {
		return map[string]interface{}{}, errors.New("error occured : " + err.Error())
	}
	response, err := post(req_url, token, "")
	if err != nil {
		return map[string]interface{}{}, errors.New("error occured : " + err.Error())
	}
	return response, nil
}

// Convert Crypto
func (c *Client) ConversionQuotes(convert_from, convert_to string) (map[string]interface{}, error) {
	req_details := ConversionQuotes{
		Convert_from: convert_from,
		Convert_to:   convert_to,
	}
	req_body := "convert_from=" + req_details.Convert_from + "&convert_to=" + req_details.Convert_to
	req_url := c.BaseURL + "/wallet/conversion-quotes"
	token, err := c.Auth()
	if err != nil {
		return map[string]interface{}{}, errors.New("error occured : " + err.Error())
	}
	response, err := post(req_url, token, req_body)
	if err != nil {
		return map[string]interface{}{}, errors.New("error occured : " + err.Error())
	}
	return response, nil
}
func (c *Client) Convert(order_id, quote_id, convert_from, convert_to, amount string) (map[string]interface{}, error) {
	req_details := Conversion{
		Order_id:     order_id,
		Quote_id:     quote_id,
		Convert_from: convert_from,
		Convert_to:   convert_to,
		Amount:       amount,
	}
	req_body := "order_id=" + req_details.Order_id + "&quote_id=" + req_details.Quote_id + "&convert_from=" + req_details.Convert_from + "&convert_to=" + req_details.Convert_to + "&amount=" + req_details.Amount
	req_url := c.BaseURL + "/wallet/convert"
	token, err := c.Auth()
	if err != nil {
		return map[string]interface{}{}, errors.New("error occured : " + err.Error())
	}
	response, err := post(req_url, token, req_body)
	if err != nil {
		return map[string]interface{}{}, errors.New("error occured : " + err.Error())
	}
	return response, nil
}
