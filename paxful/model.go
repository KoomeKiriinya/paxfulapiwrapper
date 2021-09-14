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

type InitiatePayment struct{
	Merchant string
	Apikey string 
	Apiseal string
	Nonce string
	Apisecret string
	To string
	Amount string
	Track_id string
}

type Conversion struct{
	Order_id string
	Quote_id string
	Convert_from string
	Convert_to string
	Amount string
}