package application

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ruancaetano/challenge-bravo/application/factories"
	"github.com/ruancaetano/challenge-bravo/application/handlers"
	"github.com/ruancaetano/challenge-bravo/infra/database/sqlite"
)

func Start() {
	// connect to database
	dbManager := sqlite.NewDBManager()
	err := dbManager.Open("db.sqlite")
	defer dbManager.Close()
	if err != nil {
		panic(err)
	}

	// make usecases
	getCurrencyUseCase := factories.MakeGetCurrencyUseCase(dbManager)
	addCurrencyUseCase := factories.MakeAddCurrencyUseCase(dbManager)
	convertCurrencyUseCase := factories.MakeConvertCurrencyUseCase(dbManager)
	deleteCurrencyUseCase := factories.MakeDeleteCurrencyUseCase(dbManager)

	// routes
	router := mux.NewRouter()
	router.HandleFunc("/currency/convert", handlers.MakeConvertCurrencyRequestHandler(router, convertCurrencyUseCase)).
		Methods("GET", "OPTIONS")

	router.HandleFunc("/currency/{code}", handlers.MakeGetCurrencyRequestHandler(router, getCurrencyUseCase)).
		Methods("GET", "OPTIONS")

	router.HandleFunc("/currency", handlers.MakeAddCurrencyRequestHandler(router, addCurrencyUseCase)).
		Methods("POST", "OPTIONS")

	router.HandleFunc("/currency/{code}", handlers.MakeDeleteCurrencyRequestHandler(router, deleteCurrencyUseCase)).
		Methods("DELETE", "OPTIONS")

	// bootstrap
	fmt.Println("Listen port 8000")
	err = http.ListenAndServe(":8000", router)
	if err != nil {
		panic(err)
	}
}
