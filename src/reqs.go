package src

import "fmt"

func ApiCal(fromCur string, toCur string, amount float64) {
	api_key := GetApiKey()
	fmt.Println(api_key)
}
