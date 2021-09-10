package paxful

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
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
	req.Header.Add("Authorization", apitoken)
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
