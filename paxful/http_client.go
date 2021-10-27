package paxful

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

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
