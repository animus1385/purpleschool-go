package main

import "fmt"

func main() {
	amount, from, to := inputUser()
	_ = calculate(amount, from, to)
}
func inputUser() (float64, string, string) {
	var amount float64
	var from, to string

	fmt.Scan(&amount, &from, &to)

	return amount, from, to
}

func calculate(amount float64, from string, to string) float64 {
	return 0
}
