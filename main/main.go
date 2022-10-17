package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ruancaetano/challenge-bravo/infra/database/sqlite"
	"github.com/ruancaetano/challenge-bravo/main/factories"
	"github.com/ruancaetano/challenge-bravo/main/handlers"
)

func main() {
	dbManager := sqlite.NewDBManager()
	err := dbManager.Open("db.sqlite")
	defer dbManager.Close()
	if err != nil {
		panic(err)
	}

	router := mux.NewRouter()

	// usecases
	convertCurrencyUseCase := factories.MakeConvertCurrencyUseCase(dbManager)

	router.HandleFunc("/convert", handlers.MakeConvertRequestHandler(router, convertCurrencyUseCase)).
		Methods("GET", "OPTIONS")

	err = http.ListenAndServe(":8000", router)
	if err != nil {
		panic(err)
	}
}
