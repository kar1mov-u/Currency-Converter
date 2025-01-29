package src

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/charmbracelet/huh"
	"github.com/joho/godotenv"
)

func Start() {
	var fromCur string
	huh.NewSelect[string]().
		Title("Pick your currency.").
		Options(
			huh.NewOption("USD", "USD"),
			huh.NewOption("EUR", "EUR"),
			huh.NewOption("GBP", "GBP"),
			huh.NewOption("JPY", "JPY"),
			huh.NewOption("UZS", "UZS"),
		).
		Value(&fromCur).Run()

	var toCur string
	huh.NewSelect[string]().
		Title("Pick a currency you want to convert").
		Options(
			huh.NewOption("USD", "USD"),
			huh.NewOption("EUR", "EUR"),
			huh.NewOption("GBP", "GBP"),
			huh.NewOption("JPY", "JPY"),
			huh.NewOption("UZS", "UZS"),
		).
		Value(&toCur).Run()

	var amount_input string
	huh.NewInput().
		Title("Enter amount you want to convert").
		Value(&amount_input).Run()

	amount, err := strconv.ParseFloat(amount_input, 64)
	if err != nil {
		log.Fatalf("Please enter valid numbers")
		return
	}

	res, dif := ApiCal(fromCur, toCur, amount)
	fmt.Printf("Converted: %.2f %s → %.2f %s (Rate: %.4f)\n", amount, fromCur, res, toCur, dif)

}

func GetApiKey() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .evn file %v", err)
	}
	appId := os.Getenv("APP_ID")
	return appId
}
