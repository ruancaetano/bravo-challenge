package delete_currency_test

import (
	"testing"

	"github.com/ruancaetano/challenge-bravo/internal/domain/entities"
	"github.com/ruancaetano/challenge-bravo/internal/domain/usecases/delete_currency"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type CurrencyRepositoryMock struct {
	mock.Mock
}

func (r *CurrencyRepositoryMock) Get(code string) (*entities.Currency, error) {
	args := r.Mock.Called(code)
	return args.Get(0).(*entities.Currency), args.Error(1)
}

func (r *CurrencyRepositoryMock) Create(currency *entities.Currency) error {
	args := r.Mock.Called(currency)
	return args.Error(0)
}

func (r *CurrencyRepositoryMock) Delete(code string) error {
	args := r.Mock.Called(code)
	return args.Error(0)
}

func TestDeleteCurrencyUseCase_Execute(t *testing.T) {
	currencyRepositoryMock := &CurrencyRepositoryMock{}

	currencyRepositoryMock.On("Delete", mock.Anything).Return(nil)

	usecase := delete_currency.NewDeleteCurrencyUseCase(currencyRepositoryMock)

	err := usecase.Execute(&delete_currency.DeleteCurrencyUseCaseInputDTO{
		Code: "Anything",
	})

	assert.Nil(t, err)

	currencyRepositoryMock.AssertExpectations(t)
	currencyRepositoryMock.AssertNumberOfCalls(t, "Delete", 1)
}
