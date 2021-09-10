package paxful

type Auth struct{
	Client_id string
	Client_secret string
	Grant_type string
}

type Transactions struct{
	Page string
	Limit string
	Type string
	Crypto_currency_code string
}