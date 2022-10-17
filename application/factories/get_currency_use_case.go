package factories

import (
	"github.com/ruancaetano/challenge-bravo/domain/usecases/delete_currency"
	"github.com/ruancaetano/challenge-bravo/infra/database/sqlite"
	"github.com/ruancaetano/challenge-bravo/infra/database/sqlite/repositories"
)

func MakeDeleteCurrencyUseCase(dbManager *sqlite.DBManager) *delete_currency.DeleteCurrencyUseCase {
	repository := repositories.NewSqliteCurrencyRepository(dbManager)

	return delete_currency.NewDeleteCurrencyUseCase(repository)
}
