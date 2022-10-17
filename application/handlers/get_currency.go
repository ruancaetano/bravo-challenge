package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ruancaetano/challenge-bravo/domain/usecases/get_currency"
)

func MakeGetCurrencyRequestHandler(r *mux.Router, usecase *get_currency.GetCurrencyUseCase) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		vars := mux.Vars(r)

		code := vars["code"]

		output, err := usecase.Execute(&get_currency.GetCurrencyUseCaseInputDTO{
			Code: code,
		})
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
