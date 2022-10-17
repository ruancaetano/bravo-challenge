package factories

import (
	"github.com/ruancaetano/challenge-bravo/domain/usecases/add_currency"
	"github.com/ruancaetano/challenge-bravo/infra/database/sqlite"
	"github.com/ruancaetano/challenge-bravo/infra/database/sqlite/repositories"
)

func MakeAddCurrencyUseCase(dbManager *sqlite.DBManager) *add_currency.AddCurrencyUseCase {
	repository := repositories.NewSqliteCurrencyRepository(dbManager)

	return add_currency.NewAddCurrencyUseCase(repository)
}
