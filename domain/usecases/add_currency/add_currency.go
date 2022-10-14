package add_currency

import (
	"github.com/ruancaetano/challenge-bravo/domain/entities"
	"github.com/ruancaetano/challenge-bravo/domain/repositories"
)

type AddCurrencyUseCaseInputDTO struct {
	Code                  string
	Type                  entities.CurrencyType
	DollarBasedProportion float64
}

type AddCurrencyUseCaseOutputDTO struct {
	ID                    string
	Code                  string
	Type                  entities.CurrencyType
	DollarBasedProportion float64
}

type AddCurrencyUseCase struct {
	CurrencyRepository repositories.CurrencyRepository
}

func NewAddCurrencyUseCase(currencyRepository repositories.CurrencyRepository) *AddCurrencyUseCase {
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
		Code:                  currency.Code,
		Type:                  currency.Type,
		DollarBasedProportion: currency.DollarBasedProportion,
	}, nil
}
