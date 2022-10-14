package delete_currency

import (
	"github.com/ruancaetano/challenge-bravo/domain/repositories"
)

type DeleteCurrencyUseCaseInputDTO struct {
	Code string
}

type GetCurrencyUseCase struct {
	CurrencyRepository repositories.CurrencyRepositoryInterface
}

func NewDeleteCurrencyUseCase(currencyRepository repositories.CurrencyRepositoryInterface) *GetCurrencyUseCase {
	return &GetCurrencyUseCase{
		CurrencyRepository: currencyRepository,
	}
}

func (u *GetCurrencyUseCase) Execute(input *DeleteCurrencyUseCaseInputDTO) error {
	err := u.CurrencyRepository.Delete(input.Code)
	if err != nil {
		return err
	}

	return nil
}
