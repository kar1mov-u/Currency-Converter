package src

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Data map[string]struct {
		Code  string  `json:"code"`
		Value float64 `json:"value"`
	}
}

func ApiCal(fromCur string, toCur string, amount float64) (float64, float64) {
	var data Response
	// var selected_data Data
	api_key := GetApiKey()

	URL := "https://api.currencyapi.com/v3/latest?apikey=" + api_key + "&currencies=" + toCur + "&base_currency=" + fromCur

	resp, err := http.Get(URL)
	if err != nil {
		log.Fatalf("Failed to send a request %v", err)
	}

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Fatalf("Failed to parse JSON %v", err)
	}
	return (data.Data[toCur].Value * amount), data.Data[toCur].Value

}
