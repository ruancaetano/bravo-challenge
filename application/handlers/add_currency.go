package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ruancaetano/challenge-bravo/domain/usecases/add_currency"
)

func MakeAddCurrencyRequestHandler(r *mux.Router, usecase *add_currency.AddCurrencyUseCase) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		inputDto := &add_currency.AddCurrencyUseCaseInputDTO{}

		err := json.NewDecoder(r.Body).Decode(inputDto)
		if err != nil {
			HandleError(w, err.Error())
			return
		}

		output, err := usecase.Execute(inputDto)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			HandleError(w, err.Error())
			return
		}

		responsebody, err := json.Marshal(output)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		w.Write(responsebody)
	}
}
