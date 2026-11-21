package main

import "fmt"

func main() {
	const USD_EUR = 0.9
	const USD_RUB = 81

	EUR_RUB := USD_RUB / USD_EUR

	fmt.Println("result =", EUR_RUB)
}