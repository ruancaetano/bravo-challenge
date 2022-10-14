package get_currency_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/ruancaetano/challenge-bravo/domain/entities"
	"github.com/ruancaetano/challenge-bravo/domain/usecases/get_currency"
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

func TestGetCurrencyUseCase_Execute(t *testing.T) {
	currencyRepositoryMock := &CurrencyRepositoryMock{}

	currencyRepositoryMock.On("Get", mock.Anything).Return(&entities.Currency{
		ID:   uuid.NewString(),
		Code: "Anything",
	}, nil)

	usecase := get_currency.NewGetCurrencyUseCase(currencyRepositoryMock)

	output, err := usecase.Execute(&get_currency.GetCurrencyUseCaseInputDTO{
		Code: "Anything",
	})

	assert.Nil(t, err)
	assert.NotEmpty(t, output.ID)
	assert.Equal(t, output.Code, "Anything")

	currencyRepositoryMock.AssertExpectations(t)
	currencyRepositoryMock.AssertNumberOfCalls(t, "Get", 1)
}
