package factories

import (
	"github.com/ruancaetano/challenge-bravo/domain/usecases/get_currency"
	"github.com/ruancaetano/challenge-bravo/infra/database/sqlite"
	"github.com/ruancaetano/challenge-bravo/infra/database/sqlite/repositories"
)

func MakeGetCurrencyUseCase(dbManager *sqlite.DBManager) *get_currency.GetCurrencyUseCase {
	repository := repositories.NewSqliteCurrencyRepository(dbManager)

	return get_currency.NewGetCurrencyUseCase(repository)
}
