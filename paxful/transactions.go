package paxful

import "errors"

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
