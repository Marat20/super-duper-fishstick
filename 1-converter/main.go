package main

import "fmt"

func scanCurrency() float64 {
	var currency float64
	fmt.Scan(&currency)

	return currency
}

func converter(num float64, currencyFrom, currencyTo string) {}

func main() {
	const USD_EUR = 0.9
	const USD_RUB = 81

	EUR_RUB := USD_RUB / USD_EUR

	fmt.Println("result =", EUR_RUB)

	res := scanCurrency()
	converter(res, "RUB", "EUR")
}
