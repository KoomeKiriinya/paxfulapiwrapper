package paxful

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"time"
)

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
