package entities_test

import (
	"testing"

	"github.com/ruancaetano/challenge-bravo/domain/entities"
	"github.com/stretchr/testify/assert"
)

func TestCreateNewCurrency(t *testing.T) {
	currencyFiat, err := entities.NewCurrency("USD", entities.FIAT, 0)

	assert.Nil(t, err)
	assert.NotEmpty(t, currencyFiat.ID)

	currencyCrypto, err := entities.NewCurrency("BTC", entities.CRYPTO, 0)

	assert.Nil(t, err)
	assert.NotEmpty(t, currencyCrypto.ID)

	currencyFicticious, err := entities.NewCurrency("BTC", entities.FICTICIOUS, float64(0.1))

	assert.Nil(t, err)
	assert.NotEmpty(t, currencyFicticious.ID)
}

func TestCreateNewCurrencyWithIvalidType(t *testing.T) {
	_, err := entities.NewCurrency("USD", "invalid", 0)

	assert.Error(t, err, "invalid currency type")
}

func TestCreateNewFicticiouCurrencyWithInvalidBaseValue(t *testing.T) {
	_, err := entities.NewCurrency("USD", entities.FICTICIOUS, 0)

	assert.Error(t, err, "dollar based proportion must be greater than zero")

	_, err = entities.NewCurrency("USD", entities.FICTICIOUS, -10)

	assert.Error(t, err, "dollar based proportion must be greater than zero")
}
