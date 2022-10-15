package repositories_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ruancaetano/challenge-bravo/domain/entities"
	"github.com/ruancaetano/challenge-bravo/infra/database/sqlite"
	"github.com/ruancaetano/challenge-bravo/infra/database/sqlite/repositories"
)

var manager = sqlite.NewDBManager()

func setUp() *repositories.SqliteCurrencyRepository {
	manager.Open(":memory:")

	craeteTable(manager.DB)

	repository := repositories.NewSqliteCurrencyRepository(manager)

	return repository
}

func craeteTable(db *sql.DB) {
	queryString := `CREATE TABLE currencies (
		"id" string,
		"code" string unique,
		"type" string,
		"dollar_based_proportion" float
	);`

	stmt, err := db.Prepare(queryString)
	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func TestSqliteCurrencyRepository_Create(t *testing.T) {
	t.Run("With Valid Currency", func(t *testing.T) {
		repository := setUp()
		defer repository.DBManager.Close()

		currencyFiat, err := entities.NewCurrency("USD", entities.FIAT, 0)

		assert.Nil(t, err)

		err = repository.Create(currencyFiat)

		assert.Nil(t, err)
	})

	t.Run("With Duplicated Code", func(t *testing.T) {
		repository := setUp()
		defer repository.DBManager.Close()

		currencyFiat, err := entities.NewCurrency("USD", entities.FIAT, 0)
		assert.Nil(t, err)

		err = repository.Create(currencyFiat)
		assert.Nil(t, err)

		err = repository.Create(currencyFiat)
		assert.NotNil(t, err)
	})
}

func TestSqliteCurrencyRepository_Get(t *testing.T) {
	repository := setUp()
	defer repository.DBManager.Close()

	currency, err := entities.NewCurrency("USD", entities.FIAT, 0)
	assert.Nil(t, err)

	err = repository.Create(currency)
	assert.Nil(t, err)

	foundCurrency, err := repository.Get(currency.Code)
	assert.Nil(t, err)
	assert.Equal(t, foundCurrency.ID, currency.ID)
}

func TestSqliteCurrencyRepository_GetWithInvalidToken(t *testing.T) {
	repository := setUp()
	defer repository.DBManager.Close()

	_, err := repository.Get("invalid")
	assert.NotNil(t, err)
}

func TestSqliteCurrencyRepository_Delete(t *testing.T) {
	repository := setUp()
	defer repository.DBManager.Close()

	currency, err := entities.NewCurrency("USD", entities.FIAT, 0)
	assert.Nil(t, err)

	err = repository.Create(currency)
	assert.Nil(t, err)

	err = repository.Delete(currency.Code)
	assert.Nil(t, err)
}

func TestSqliteCurrencyRepository_DeleteWithInvalidToken(t *testing.T) {
	repository := setUp()
	defer repository.DBManager.Close()

	err := repository.Delete("invalid")
	assert.Error(t, err, "currency not found")
}
