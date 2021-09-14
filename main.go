package main

import (
	"fmt"
	"github.com/paxfulapiwrapper/paxful"
)

func main() {
	// Initiate a client
	pax_client := paxful.NewClient()

	// list transactions 
	transactions, err := pax_client.Transactions("1", "10", "all", "USDT")
	if err != nil {
		fmt.Println("error :" + err.Error())
	}
	fmt.Println(transactions)
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
