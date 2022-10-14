package get_currency

import (
	"github.com/ruancaetano/challenge-bravo/internal/domain/entities"
	"github.com/ruancaetano/challenge-bravo/internal/domain/repositories"
)

type GetCurrencyUseCaseInputDTO struct {
	Code string
}

type GetCurrencyUseCaseOutputDTO struct {
	ID                    string
	Code                  string
	Type                  entities.CurrencyType
	DollarBasedProportion float64
}

type GetCurrencyUseCase struct {
	CurrencyRepository repositories.CurrencyRepository
}

func NewGetCurrencyUseCase(currencyRepository repositories.CurrencyRepository) *GetCurrencyUseCase {
	return &GetCurrencyUseCase{
		CurrencyRepository: currencyRepository,
	}
}

func (u *GetCurrencyUseCase) Execute(input *GetCurrencyUseCaseInputDTO) (*GetCurrencyUseCaseOutputDTO, error) {
	currency, err := u.CurrencyRepository.Get(input.Code)
	if err != nil {
		return nil, err
	}

	return &GetCurrencyUseCaseOutputDTO{
		ID:                    currency.ID,
		Code:                  currency.Code,
		Type:                  currency.Type,
		DollarBasedProportion: currency.DollarBasedProportion,
	}, nil
}
