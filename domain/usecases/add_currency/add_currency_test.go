package add_currency_test

import (
	"testing"

	"github.com/ruancaetano/challenge-bravo/domain/entities"
	"github.com/ruancaetano/challenge-bravo/domain/usecases/add_currency"
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

func TestAddCurrencyUseCase_Execute(t *testing.T) {
	currencyRepositoryMock := &CurrencyRepositoryMock{}

	currencyRepositoryMock.On("Create", mock.Anything).Return(nil)

	usecase := add_currency.NewAddCurrencyUseCase(currencyRepositoryMock)

	output, err := usecase.Execute(&add_currency.AddCurrencyUseCaseInputDTO{
		Code:                  "USD",
		Type:                  entities.FIAT,
		DollarBasedProportion: 0,
	})

	assert.Nil(t, err)
	assert.NotEmpty(t, output.ID)

	currencyRepositoryMock.AssertExpectations(t)
	currencyRepositoryMock.AssertNumberOfCalls(t, "Create", 1)
}
