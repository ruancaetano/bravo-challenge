package repositories

import (
	"errors"

	"github.com/ruancaetano/challenge-bravo/domain/entities"
	"github.com/ruancaetano/challenge-bravo/infra/database/sqlite"
)

type SqliteCurrencyRepository struct {
	DBManager *sqlite.DBManager
}

func NewSqliteCurrencyRepository(dbManager *sqlite.DBManager) *SqliteCurrencyRepository {
	return &SqliteCurrencyRepository{DBManager: dbManager}
}

func (r *SqliteCurrencyRepository) Create(currency *entities.Currency) error {
	stmt, err := r.DBManager.DB.Prepare("INSERT INTO currencies (id, code, type, dollar_based_proportion) VALUES (?,?,?,?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(currency.ID, currency.Code, currency.Type, currency.DollarBasedProportion)

	if err != nil {
		return err
	}

	return nil
}

func (r *SqliteCurrencyRepository) Get(code string) (*entities.Currency, error) {
	stmt, err := r.DBManager.DB.Prepare("SELECT id, code, type, dollar_based_proportion from currencies where code = ?")
	if err != nil {
		return nil, err
	}

	currency := &entities.Currency{}

	err = stmt.QueryRow(code).Scan(&currency.ID, &currency.Code, &currency.Type, &currency.DollarBasedProportion)

	if err != nil {
		return nil, err
	}

	return currency, nil
}

func (r *SqliteCurrencyRepository) Delete(code string) error {
	stmt, err := r.DBManager.DB.Prepare("DELETE FROM currencies where code = ?")
	if err != nil {
		return err
	}

	result, err := stmt.Exec(code)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("currency not found")
	}

	return nil
}
