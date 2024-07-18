package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {

	const (
		usdToEurRate = 0.93
		usdToRubRate = 88.50
	)

	var sourceCurrency string
	var err error

	for {
		sourceCurrency, err = choiseSourceCurrency()
		if err != nil {
			fmt.Printf("Ошибка: %s\n", err)
		} else {
			fmt.Printf("Исходная валюта: %s\n", sourceCurrency)
			break
		}
	}

	for {
		amount, err := amountCurrency()
		if err != nil {
			fmt.Println("Ошибка:", err) // Это сообщение не должно выводиться, так как функция не возвращает ошибку
		} else {
			fmt.Println("Введенная сумма: ", amount)
			break
		}

	}

}

func choiseSourceCurrency() (string, error) {
	for {
		var input string
		fmt.Print("Выберете валюту (USD, EUR, RUB): ")
		fmt.Scan(&input)
		// Проверка допустимости валюты
		if input == "USD" || input == "EUR" || input == "RUB" {
			return input, nil // Валюта допустима
		} else {
			return "", errors.New("Неправильная валюта. Попробуйте еще раз.")
			// fmt.Println("Неправильная валюта. Попробуйте еще раз.")
		}
	}
}

func amountCurrency() (float64, error) {
	for {
		var input string
		for {
			fmt.Print("Введите сумму: ")
			fmt.Scanln(&input)

			amount, err := strconv.ParseFloat(input, 64)
			if err == nil {
				return amount, nil
			} else {
				return 0, errors.New("Введенное значение не является числом. Попробуйте еще раз.")
			}

		}

	}
}
