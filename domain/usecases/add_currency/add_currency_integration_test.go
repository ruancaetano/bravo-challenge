//go:build integration
// +build integration

package add_currency_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/ruancaetano/challenge-bravo/domain/entities"
	"github.com/ruancaetano/challenge-bravo/domain/usecases/add_currency"
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

func TestAddCurrencyUseCaseIntegration_Execute(t *testing.T) {
	setUp()
	defer manager.Close()

	repository := repositories.NewSqliteCurrencyRepository(manager)

	usecase := add_currency.NewAddCurrencyUseCase(repository)

	output, err := usecase.Execute(&add_currency.AddCurrencyUseCaseInputDTO{
		Code:                  "USD",
		Type:                  entities.FIAT,
		DollarBasedProportion: 0,
	})

	assert.Nil(t, err)
	assert.NotEmpty(t, output.ID)
}
