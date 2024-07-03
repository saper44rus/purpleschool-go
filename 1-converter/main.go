package main

import "fmt"

func main() {

	const (
		usdToEurRate = 0.93
		usdToRubRate = 88.50
	)
	eurToRubRate := usdToRubRate / usdToEurRate
	fmt.Print(eurToRubRate)
}
