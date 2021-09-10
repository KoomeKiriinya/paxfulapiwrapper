package main

import (
	"fmt"

	"github.com/paxfulapiwrapper/paxful"
)

func main() {

	pax_client := paxful.NewClient()
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

}
