package main

import "github.com/ruancaetano/challenge-bravo/infra/services"

func main() {
	service := &services.AwesomeCurrencyConverterService{}

	service.Convert("USD", "BRL", 10)
}
