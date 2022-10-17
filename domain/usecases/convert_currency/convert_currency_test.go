package convert_currency_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/ruancaetano/challenge-bravo/domain/entities"
	"github.com/ruancaetano/challenge-bravo/domain/services"
	"github.com/ruancaetano/challenge-bravo/domain/usecases/convert_currency"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// CurrencyRepositoryMock
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

// CurrencyConverterServiceMock
type CurrencyConverterServiceMock struct {
	mock.Mock
}

func (r *CurrencyConverterServiceMock) Convert(fromCode string, toCode string, amount float64) (*services.CurrencyConverterServiceResponse, error) {
	args := r.Mock.Called(fromCode, toCode, amount)
	return args.Get(0).(*services.CurrencyConverterServiceResponse), args.Error(1)
}

func TestConvertCurrencyUseCase_Execute(t *testing.T) {
	t.Run("FIAT currencies convertion", func(t *testing.T) {
		repositoryMock := &CurrencyRepositoryMock{}

		repositoryMock.On("Get", "BRL").Return(&entities.Currency{
			ID:   uuid.NewString(),
			Type: entities.FIAT,
			Code: "BRL",
		}, nil)

		repositoryMock.On("Get", "USD").Return(&entities.Currency{
			ID:   uuid.NewString(),
			Type: entities.FIAT,
			Code: "USD",
		}, nil)

		serviceMock := &CurrencyConverterServiceMock{}

		serviceMock.On("Convert", "USD", "BRL", mock.Anything).Return(&services.CurrencyConverterServiceResponse{
			FromCurrencyCode: "USD",
			ToCurrencyCode:   "BRL",
			Value:            float64(100),
		}, nil)

		usecase := convert_currency.NewConvertCurrencyUseCase(repositoryMock, serviceMock)

		output, err := usecase.Execute(&convert_currency.ConvertCurrencyUseCaseInputDTO{
			FromCurrencyCode: "USD",
			ToCurrencyCode:   "BRL",
			Amount:           float64(10),
		})

		assert.Nil(t, err)
		assert.Equal(t, output.FromCurrencyCode, "USD")
		assert.Equal(t, output.ToCurrencyCode, "BRL")
		assert.Equal(t, output.Value, fmt.Sprintf("%.2f", float64(100)))
	})

	t.Run("FIAT currencies to crypto convertion", func(t *testing.T) {
		repositoryMock := &CurrencyRepositoryMock{}

		repositoryMock.On("Get", "BTC").Return(&entities.Currency{
			ID:   uuid.NewString(),
			Type: entities.CRYPTO,
			Code: "BTC",
		}, nil)

		repositoryMock.On("Get", "USD").Return(&entities.Currency{
			ID:   uuid.NewString(),
			Type: entities.FIAT,
			Code: "USD",
		}, nil)

		serviceMock := &CurrencyConverterServiceMock{}

		serviceMock.On("Convert", "USD", "BTC", mock.Anything).Return(&services.CurrencyConverterServiceResponse{
			FromCurrencyCode: "USD",
			ToCurrencyCode:   "BTC",
			Value:            float64(0.000051),
		}, nil)

		usecase := convert_currency.NewConvertCurrencyUseCase(repositoryMock, serviceMock)

		output, err := usecase.Execute(&convert_currency.ConvertCurrencyUseCaseInputDTO{
			FromCurrencyCode: "USD",
			ToCurrencyCode:   "BTC",
			Amount:           float64(10),
		})

		assert.Nil(t, err)
		assert.Equal(t, output.FromCurrencyCode, "USD")
		assert.Equal(t, output.ToCurrencyCode, "BTC")
		assert.Equal(t, output.Value, fmt.Sprintf("%.18f", float64(0.000051)))
	})

	t.Run("FIAT currencies to ficticious convertion", func(t *testing.T) {
		repositoryMock := &CurrencyRepositoryMock{}

		repositoryMock.On("Get", "BRL").Return(&entities.Currency{
			ID:   uuid.NewString(),
			Type: entities.FIAT,
			Code: "BRL",
		}, nil)

		repositoryMock.On("Get", "FIC1").Return(&entities.Currency{
			ID:                    uuid.NewString(),
			Type:                  entities.FICTICIOUS,
			Code:                  "FIC1",
			DollarBasedProportion: 0.1,
		}, nil)

		serviceMock := &CurrencyConverterServiceMock{}

		serviceMock.On("Convert", "BRL", "USD", float64(10)).Return(&services.CurrencyConverterServiceResponse{
			FromCurrencyCode: "USD",
			ToCurrencyCode:   "BRL",
			Value:            float64(50),
		}, nil)

		usecase := convert_currency.NewConvertCurrencyUseCase(repositoryMock, serviceMock)

		output, err := usecase.Execute(&convert_currency.ConvertCurrencyUseCaseInputDTO{
			FromCurrencyCode: "BRL",
			ToCurrencyCode:   "FIC1",
			Amount:           float64(10),
		})

		assert.Nil(t, err)
		assert.Equal(t, output.FromCurrencyCode, "BRL")
		assert.Equal(t, output.ToCurrencyCode, "FIC1")
		assert.Equal(t, output.Value, fmt.Sprintf("%.2f", float64(50*0.1)))
	})

	t.Run("Ficticious currencies convertion", func(t *testing.T) {
		repositoryMock := &CurrencyRepositoryMock{}

		repositoryMock.On("Get", "FIC1").Return(&entities.Currency{
			ID:                    uuid.NewString(),
			Type:                  entities.FICTICIOUS,
			Code:                  "FIC1",
			DollarBasedProportion: 1,
		}, nil)

		repositoryMock.On("Get", "FIC2").Return(&entities.Currency{
			ID:                    uuid.NewString(),
			Type:                  entities.FICTICIOUS,
			Code:                  "FIC2",
			DollarBasedProportion: 5,
		}, nil)

		serviceMock := &CurrencyConverterServiceMock{}

		usecase := convert_currency.NewConvertCurrencyUseCase(repositoryMock, serviceMock)

		output, err := usecase.Execute(&convert_currency.ConvertCurrencyUseCaseInputDTO{
			FromCurrencyCode: "FIC1",
			ToCurrencyCode:   "FIC2",
			Amount:           float64(10),
		})

		assert.Nil(t, err)
		assert.Equal(t, output.FromCurrencyCode, "FIC1")
		assert.Equal(t, output.ToCurrencyCode, "FIC2")
		assert.Equal(t, output.Value, fmt.Sprintf("%.2f", float64(5*1*10)))
	})

	t.Run("Ficticious currencies to FIAT convertion", func(t *testing.T) {
		repositoryMock := &CurrencyRepositoryMock{}

		repositoryMock.On("Get", "FIC1").Return(&entities.Currency{
			ID:                    uuid.NewString(),
			Type:                  entities.FICTICIOUS,
			Code:                  "FIC1",
			DollarBasedProportion: 5,
		}, nil)

		repositoryMock.On("Get", "BRL").Return(&entities.Currency{
			ID:   uuid.NewString(),
			Type: entities.FIAT,
			Code: "BRL",
		}, nil)

		serviceMock := &CurrencyConverterServiceMock{}

		serviceMock.On("Convert", "USD", "BRL", float64(50)).Return(&services.CurrencyConverterServiceResponse{
			FromCurrencyCode: "USD",
			ToCurrencyCode:   "BRL",
			Value:            float64(250),
		}, nil)

		usecase := convert_currency.NewConvertCurrencyUseCase(repositoryMock, serviceMock)

		output, err := usecase.Execute(&convert_currency.ConvertCurrencyUseCaseInputDTO{
			FromCurrencyCode: "FIC1",
			ToCurrencyCode:   "BRL",
			Amount:           float64(10),
		})

		assert.Nil(t, err)
		assert.Equal(t, output.FromCurrencyCode, "FIC1")
		assert.Equal(t, output.ToCurrencyCode, "BRL")
		assert.Equal(t, output.Value, fmt.Sprintf("%.2f", float64(5*10*5)))
	})

	t.Run("Currency not supported", func(t *testing.T) {
		repositoryMock := &CurrencyRepositoryMock{}

		repositoryMock.On("Get", mock.Anything).Return(&entities.Currency{}, errors.New("currency not supported"))

		serviceMock := &CurrencyConverterServiceMock{}

		usecase := convert_currency.NewConvertCurrencyUseCase(repositoryMock, serviceMock)

		_, err := usecase.Execute(&convert_currency.ConvertCurrencyUseCaseInputDTO{
			FromCurrencyCode: "INVALID1",
			ToCurrencyCode:   "INVALID2",
			Amount:           float64(10),
		})

		assert.Error(t, err, "currency not supported")
	})
}
