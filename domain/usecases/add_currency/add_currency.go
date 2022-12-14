package add_currency

import (
	"strings"

	"github.com/ruancaetano/challenge-bravo/domain/entities"
	"github.com/ruancaetano/challenge-bravo/domain/repositories"
)

type AddCurrencyUseCaseInputDTO struct {
	Code                  string                `json:"code"`
	Type                  entities.CurrencyType `json:"type"`
	DollarBasedProportion float64               `json:"dollarBasedProportion"`
}

type AddCurrencyUseCaseOutputDTO struct {
	ID                    string                `json:"id"`
	Code                  string                `json:"code"`
	Type                  entities.CurrencyType `json:"type"`
	DollarBasedProportion float64               `json:"dollarBasedProportion"`
}

type AddCurrencyUseCase struct {
	CurrencyRepository repositories.CurrencyRepositoryInterface
}

func NewAddCurrencyUseCase(currencyRepository repositories.CurrencyRepositoryInterface) *AddCurrencyUseCase {
	return &AddCurrencyUseCase{
		CurrencyRepository: currencyRepository,
	}
}

func (u *AddCurrencyUseCase) Execute(input *AddCurrencyUseCaseInputDTO) (*AddCurrencyUseCaseOutputDTO, error) {
	currency, err := entities.NewCurrency(input.Code, input.Type, input.DollarBasedProportion)
	if err != nil {
		return nil, err
	}

	err = u.CurrencyRepository.Create(currency)

	if err != nil {
		return nil, err
	}

	return &AddCurrencyUseCaseOutputDTO{
		ID:                    currency.ID,
		Code:                  strings.ToUpper(currency.Code),
		Type:                  currency.Type,
		DollarBasedProportion: currency.DollarBasedProportion,
	}, nil
}
