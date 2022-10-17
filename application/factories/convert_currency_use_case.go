package factories

import (
	"github.com/ruancaetano/challenge-bravo/domain/usecases/convert_currency"
	"github.com/ruancaetano/challenge-bravo/infra/cache"
	"github.com/ruancaetano/challenge-bravo/infra/database/sqlite"
	"github.com/ruancaetano/challenge-bravo/infra/database/sqlite/repositories"
	"github.com/ruancaetano/challenge-bravo/infra/services"
)

func MakeConvertCurrencyUseCase(dbManager *sqlite.DBManager) *convert_currency.ConvertCurrencyUseCase {
	repository := repositories.NewSqliteCurrencyRepository(dbManager)
	memoryCache := cache.NewMemoryCache()
	service := services.NewAwesomeCurrencyConverterService(memoryCache)

	return convert_currency.NewConvertCurrencyUseCase(repository, service)
}
