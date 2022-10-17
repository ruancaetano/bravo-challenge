package services

type CurrencyConverterServiceResponse struct {
	FromCurrencyCode string
	ToCurrencyCode   string
	Value            float64
}

type CurrencyConverterServiceInterface interface {
	Convert(fromCode string, toCode string, amount float64) (*CurrencyConverterServiceResponse, error)
}
