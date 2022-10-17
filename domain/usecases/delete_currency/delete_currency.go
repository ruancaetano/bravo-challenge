package delete_currency

import (
	"github.com/ruancaetano/challenge-bravo/domain/repositories"
)

type DeleteCurrencyUseCaseInputDTO struct {
	Code string `json:"code"`
}

type DeleteCurrencyUseCase struct {
	CurrencyRepository repositories.CurrencyRepositoryInterface
}

func NewDeleteCurrencyUseCase(currencyRepository repositories.CurrencyRepositoryInterface) *DeleteCurrencyUseCase {
	return &DeleteCurrencyUseCase{
		CurrencyRepository: currencyRepository,
	}
}

func (u *DeleteCurrencyUseCase) Execute(input *DeleteCurrencyUseCaseInputDTO) error {
	err := u.CurrencyRepository.Delete(input.Code)
	if err != nil {
		return err
	}

	return nil
}
