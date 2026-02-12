package main

import "fmt"

func main() {
	const USD_TO_EUR = 0.84
	const USD_TO_RUB = 77.19

	const EUR_TO_RUB = USD_TO_RUB / USD_TO_EUR

	fmt.Println(EUR_TO_RUB)
}