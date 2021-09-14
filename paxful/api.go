package paxful

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
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

// list all transactions
func (c *Client) Transactions(page, limit, trans_type, code string) (map[string]interface{}, error) {
	req_details := Transactions{
		Page:                 page,
		Limit:                limit,
		Type:                 trans_type,
		Crypto_currency_code: code,
	}
	req_body := "page=" + req_details.Page + "&limit=" + req_details.Limit + "&type=" + req_details.Type + "&crypto_currency_code=" + req_details.Crypto_currency_code
	req_url := c.BaseURL + "/transactions/all"
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
	req_url := c.BaseURL + "/wallet/converst"
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

func post(apiurl string, apitoken string, body string) (map[string]interface{}, error) {
	payload := strings.NewReader(body)
	client := &http.Client{}
	req, err := http.NewRequest("POST", apiurl, payload)
	if err != nil {
		return map[string]interface{}{}, errors.New("error occured : " + err.Error())
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	// due to requests that dont require api token
	if apitoken != "" {
		req.Header.Add("Authorization", "Bearer "+apitoken)
	}

	res, err := client.Do(req)
	if err != nil {
		return map[string]interface{}{}, errors.New("error occured : " + err.Error())
	}

	defer res.Body.Close()
	res_body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return map[string]interface{}{}, errors.New("error occured : " + err.Error())
	}

	var response map[string]interface{}
	json.Unmarshal([]byte(res_body), &response)
	return response, nil

}

func get(apiurl string, apitoken string) (map[string]interface{}, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", apiurl, nil)
	if err != nil {
		return map[string]interface{}{}, errors.New("error occured : " + err.Error())
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Bearer "+apitoken)
	res, err := client.Do(req)
	if err != nil {
		return map[string]interface{}{}, errors.New("error occured : " + err.Error())
	}

	defer res.Body.Close()
	res_body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return map[string]interface{}{}, errors.New("error occured : " + err.Error())
	}

	var response map[string]interface{}
	json.Unmarshal([]byte(res_body), &response)
	return response, nil

}

func InitiatePaxfulPayment(amount, track_id string) string {
	payment_details := InitiatePayment{
		Merchant:  os.Getenv("MERCHANT_ID"),
		Apikey:    os.Getenv("API_KEY"),
		Apisecret: os.Getenv("API_SECRET"),
		Nonce:     fmt.Sprint(time.Now().Unix()),
		To:        os.Getenv("BITCOIN_ADDRESS"),
		Amount:    amount,
		Track_id:  track_id,
	}

	payments_api_seal_body := "merchant=" + payment_details.Merchant + "&apikey=" + payment_details.Apikey + "&nonce=" + payment_details.Nonce + "&to=" + payment_details.To + "&track_id=" + payment_details.Track_id + "&amount=" + payment_details.Amount
	h := hmac.New(sha256.New, []byte(payment_details.Apisecret))
	h.Write([]byte(payments_api_seal_body))
	api_seal := hex.EncodeToString(h.Sum(nil))
	payment_link := "https://paxful.com/wallet/pay?" + payments_api_seal_body + "&apiseal=" + api_seal
	return payment_link

}
