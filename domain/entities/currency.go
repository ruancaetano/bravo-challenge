package entities

import (
	"errors"

	"github.com/google/uuid"
)

type CurrencyType string

const (
	FIAT       CurrencyType = "FIAT"
	CRYPTO     CurrencyType = "CRYPTO"
	FICTICIOUS CurrencyType = "FICTICIOUS"
)

type Currency struct {
	ID                    string       `json:"id"`
	Code                  string       `json:"code"`
	Type                  CurrencyType `json:"type"`
	DollarBasedProportion float64      `json:"dollarBasedProportion"`
}

func NewCurrency(code string, currencyType CurrencyType, dollarBasedProportion float64) (*Currency, error) {
	currency := &Currency{
		ID:                    uuid.NewString(),
		Code:                  code,
		Type:                  currencyType,
		DollarBasedProportion: dollarBasedProportion,
	}

	err := currency.Validate()
	if err != nil {
		return &Currency{}, err
	}

	return currency, nil
}

func (c *Currency) Validate() error {
	if c.Type != FIAT && c.Type != CRYPTO && c.Type != FICTICIOUS {
		return errors.New("invalid currency type")
	}

	if c.Type == FICTICIOUS && c.DollarBasedProportion <= 0 {
		return errors.New("dollar based proportion must be greater than zero")
	}

	return nil
}
