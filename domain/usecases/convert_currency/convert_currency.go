package convert_currency

import (
	"fmt"
	"strings"

	"github.com/ruancaetano/challenge-bravo/domain/entities"
	"github.com/ruancaetano/challenge-bravo/domain/repositories"
	"github.com/ruancaetano/challenge-bravo/domain/services"
)

type ConvertCurrencyUseCaseInputDTO struct {
	FromCurrencyCode string
	ToCurrencyCode   string
	Amount           float64
}

type ConvertCurrencyUseCaseOutputDTO struct {
	FromCurrencyCode string
	ToCurrencyCode   string
	Value            string
}

type ConvertCurrencyUseCase struct {
	CurrencyRepository       repositories.CurrencyRepositoryInterface
	CurrencyConverterService services.CurrencyConverterServiceInterface
}

func NewConvertCurrencyUseCase(repository repositories.CurrencyRepositoryInterface, service services.CurrencyConverterServiceInterface) *ConvertCurrencyUseCase {
	return &ConvertCurrencyUseCase{
		CurrencyRepository:       repository,
		CurrencyConverterService: service,
	}
}

func (u *ConvertCurrencyUseCase) Execute(input *ConvertCurrencyUseCaseInputDTO) (*ConvertCurrencyUseCaseOutputDTO, error) {
	fromCurrency, err := u.CurrencyRepository.Get(strings.ToUpper(input.FromCurrencyCode))
	if err != nil {
		return nil, fmt.Errorf("%s is not a supported currency", strings.ToUpper(input.FromCurrencyCode))
	}

	toCurrency, err := u.CurrencyRepository.Get(strings.ToUpper(input.ToCurrencyCode))
	if err != nil {
		return nil, fmt.Errorf("%s is not a supported currency", strings.ToUpper(input.ToCurrencyCode))
	}

	if fromCurrency.Type != entities.FICTICIOUS && toCurrency.Type != entities.FICTICIOUS {
		return u.convertRealCurrencies(fromCurrency, toCurrency, input.Amount)
	}

	if fromCurrency.Type != entities.FICTICIOUS && toCurrency.Type == entities.FICTICIOUS {
		return u.convertFromRealToFicticiousCurrency(fromCurrency, toCurrency, input.Amount)
	}

	if fromCurrency.Type == entities.FICTICIOUS && toCurrency.Type != entities.FICTICIOUS {
		return u.convertFromFicticiousToRealCurrency(fromCurrency, toCurrency, input.Amount)
	}

	return u.convertFicticiousCurrencies(fromCurrency, toCurrency, input.Amount)
}

func (u *ConvertCurrencyUseCase) convertRealCurrencies(fromCurrency *entities.Currency, toCurrency *entities.Currency, amount float64) (*ConvertCurrencyUseCaseOutputDTO, error) {
	result, err := u.CurrencyConverterService.Convert(fromCurrency.Code, toCurrency.Code, amount)
	if err != nil {
		return nil, err
	}

	return &ConvertCurrencyUseCaseOutputDTO{
		FromCurrencyCode: fromCurrency.Code,
		ToCurrencyCode:   toCurrency.Code,
		Value:            formatValueBasedOnCurrency(toCurrency, result.Value),
	}, nil
}

func (u *ConvertCurrencyUseCase) convertFromRealToFicticiousCurrency(fromCurrency *entities.Currency, toCurrency *entities.Currency, amount float64) (*ConvertCurrencyUseCaseOutputDTO, error) {
	result, err := u.CurrencyConverterService.Convert(fromCurrency.Code, "USD", amount)
	if err != nil {
		return nil, err
	}

	return &ConvertCurrencyUseCaseOutputDTO{
		FromCurrencyCode: fromCurrency.Code,
		ToCurrencyCode:   toCurrency.Code,
		Value:            formatValueBasedOnCurrency(toCurrency, (result.Value * toCurrency.DollarBasedProportion)),
	}, nil
}

func (u *ConvertCurrencyUseCase) convertFromFicticiousToRealCurrency(fromCurrency *entities.Currency, toCurrency *entities.Currency, amount float64) (*ConvertCurrencyUseCaseOutputDTO, error) {
	ficticiousAmountInDollar := amount * fromCurrency.DollarBasedProportion

	result, err := u.CurrencyConverterService.Convert("USD", toCurrency.Code, ficticiousAmountInDollar)
	if err != nil {
		return nil, err
	}

	return &ConvertCurrencyUseCaseOutputDTO{
		FromCurrencyCode: fromCurrency.Code,
		ToCurrencyCode:   toCurrency.Code,
		Value:            formatValueBasedOnCurrency(toCurrency, result.Value),
	}, nil
}

func (u *ConvertCurrencyUseCase) convertFicticiousCurrencies(fromCurrency *entities.Currency, toCurrency *entities.Currency, amount float64) (*ConvertCurrencyUseCaseOutputDTO, error) {
	ficticiousAmountInDollar := amount * fromCurrency.DollarBasedProportion

	return &ConvertCurrencyUseCaseOutputDTO{
		FromCurrencyCode: fromCurrency.Code,
		ToCurrencyCode:   toCurrency.Code,
		Value:            formatValueBasedOnCurrency(toCurrency, (ficticiousAmountInDollar * toCurrency.DollarBasedProportion)),
	}, nil
}

func formatValueBasedOnCurrency(currency *entities.Currency, value float64) string {
	if currency.Type == entities.CRYPTO {
		return fmt.Sprintf("%.18f", value)
	}
	return fmt.Sprintf("%.2f", value)
}
