package main

import "fmt"

func main() {

	const (
		usdToEurRate = 0.93
		usdToRubRate = 88.50
	)
	eurToRubRate := usdToRubRate / usdToEurRate
	fmt.Println(eurToRubRate)
	fmt.Println(readUserParams())
}

func readUserParams() string {
    var userData string
    fmt.Print("Enter your data: ")
    fmt.Scanln(&userData)
	return userData
}

func currencyCalculation(summCurrency int,sourceCurrency, targetCurrency string) float64 {

}