package delete_currency_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/ruancaetano/challenge-bravo/domain/entities"
	"github.com/ruancaetano/challenge-bravo/domain/usecases/delete_currency"
	"github.com/ruancaetano/challenge-bravo/infra/database/sqlite"
	"github.com/ruancaetano/challenge-bravo/infra/database/sqlite/repositories"
	"github.com/stretchr/testify/assert"
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

func TestDeleteCurrencyUseCaseIntegration_Execute(t *testing.T) {
	repository := setUp()
	defer manager.Close()

	usecase := delete_currency.NewDeleteCurrencyUseCase(repository)

	currency, err := entities.NewCurrency("USD", entities.FIAT, 0)
	assert.Nil(t, err)

	err = repository.Create(currency)
	assert.Nil(t, err)

	err = usecase.Execute(&delete_currency.DeleteCurrencyUseCaseInputDTO{
		Code: "USD",
	})

	assert.Nil(t, err)

	_, err = repository.Get("USD")
	assert.Error(t, err, "sql: no rows in result set")
}
