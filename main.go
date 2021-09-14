package main

import (
	"fmt"
	"time"

	"github.com/paxfulapiwrapper/paxful"
)

func main() {

	pax_client := paxful.NewClient()
	transactions, err := pax_client.Transactions("1", "10", "all", "USDT")
	if err != nil {
		fmt.Println("error :" + err.Error())
	}
	transactions_value := transactions["data"].(map[string]interface{})["transactions"]

	for k, v := range transactions_value.([]interface{}) {
		fmt.Println(v.(map[string]interface{})["amount_fiat"])
		fmt.Println(k)

	}
	fmt.Println(transactions_value.([]interface{})[0])
	fmt.Println(time.Now().Unix())
	// values of currencies per country ie. USDT to KES BTC to KES etc
	currencies_list, err := pax_client.CurrencyList()
	if err != nil {
		fmt.Println("error :" + err.Error())
	}
	fmt.Println(currencies_list)
	// values of BTC and USD per country
	currency_rates, err := pax_client.CurrencyRates()
	if err != nil {
		fmt.Println("error :" + err.Error())
	}
	fmt.Println(currency_rates)
	// Initiate a BTC payment link
	payment_link := paxful.InitiatePaxfulPayment("0.000225", "xx12")
	fmt.Println(payment_link)
}
