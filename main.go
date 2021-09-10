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

}
