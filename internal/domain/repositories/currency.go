package repositories

import "github.com/ruancaetano/challenge-bravo/internal/domain/entities"

type CurrencyRepository interface {
	Get(code string) (*entities.Currency, error)
	Create(*entities.Currency) error
	Delete(code string) error
}
