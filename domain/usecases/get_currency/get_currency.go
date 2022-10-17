package get_currency

import (
	"strings"

	"github.com/ruancaetano/challenge-bravo/domain/entities"
	"github.com/ruancaetano/challenge-bravo/domain/repositories"
)

type GetCurrencyUseCaseInputDTO struct {
	Code string `json:"code"`
}

type GetCurrencyUseCaseOutputDTO struct {
	ID                    string                `json:"id"`
	Code                  string                `json:"code"`
	Type                  entities.CurrencyType `json:"type"`
	DollarBasedProportion float64               `json:"dollarBasedProportion"`
}

type GetCurrencyUseCase struct {
	CurrencyRepository repositories.CurrencyRepositoryInterface
}

func NewGetCurrencyUseCase(currencyRepository repositories.CurrencyRepositoryInterface) *GetCurrencyUseCase {
	return &GetCurrencyUseCase{
		CurrencyRepository: currencyRepository,
	}
}

func (u *GetCurrencyUseCase) Execute(input *GetCurrencyUseCaseInputDTO) (*GetCurrencyUseCaseOutputDTO, error) {
	currency, err := u.CurrencyRepository.Get(strings.ToUpper(input.Code))
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
