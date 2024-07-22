package main

import (
	"errors"
	"fmt"
	"strconv"
)

const (
	usdToEurRate = 0.93
	usdToRubRate = 88.50
)

func main() {
	sourceCurrency := getSourceCurrency()
	amount := getAmount()
	targetCurrency := getTargetCurrency(sourceCurrency)

	convertedAmount, err := convertCurrency(amount, sourceCurrency, targetCurrency)
	if err != nil {
		fmt.Println("Ошибка конвертации:", err)
	} else {
		fmt.Printf("%.2f %s = %.2f %s\n", amount, sourceCurrency, convertedAmount, targetCurrency)
	}
}

func convertCurrency(amount float64, sourceCurrency, targetCurrency string) (float64, error) {

	// Все конвертации проходят через доллар США
	var usdAmount float64
	switch sourceCurrency {
	case "USD":
		usdAmount = amount
	case "EUR":
		usdAmount = amount / usdToEurRate
	case "RUB":
		usdAmount = amount / usdToRubRate
	default:
		return 0, errors.New("Неподдерживаемая исходная валюта")
	}

	switch targetCurrency {
	case "USD":
		return usdAmount, nil
	case "EUR":
		return usdAmount * usdToEurRate, nil
	case "RUB":
		return usdAmount * usdToRubRate, nil
	default:
		return 0, errors.New("Неподдерживаемая целевая валюта")
	}

}

func getSourceCurrency() string {
	for {
		sourceCurrency, err := choiseSourceCurrency()
		if err != nil {
			fmt.Printf("Ошибка: %s\n", err)
		} else {
			return sourceCurrency
		}
	}
}

func getAmount() float64 {
	for {
		amount, err := amountCurrency()
		if err != nil {
			fmt.Println("Ошибка:", err)
		} else {
			return amount
		}
	}
}

func getTargetCurrency(sourceCurrency string) string {
	for {
		targetCurrency, err := targetCurrency(sourceCurrency)
		if err != nil {
			fmt.Println("Ошибка:", err)
		} else {
			return targetCurrency
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

func targetCurrency(sourceCurrency string) (string, error) {
	for {
		var prompt string
		switch sourceCurrency {
		case "USD":
			prompt = "Выберите целевую валюту (EUR, RUB): "
		case "EUR":
			prompt = "Выберите целевую валюту (USD, RUB): "
		case "RUB":
			prompt = "Выберите целевую валюту (USD, EUR): "
		}
		fmt.Print(prompt)
		var input string
		fmt.Scan(&input)

		// Проверка введенной валюты
		if (sourceCurrency == "USD" && (input == "EUR" || input == "RUB")) ||
			(sourceCurrency == "EUR" && (input == "USD" || input == "RUB")) ||
			(sourceCurrency == "RUB" && (input == "USD" || input == "EUR")) {
			return input, nil
		}

		return "", errors.New("Неправильная валюта. Попробуйте еще раз.")
	}
}
