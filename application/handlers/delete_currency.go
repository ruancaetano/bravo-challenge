package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ruancaetano/challenge-bravo/domain/usecases/delete_currency"
)

func MakeDeleteCurrencyRequestHandler(r *mux.Router, usecase *delete_currency.DeleteCurrencyUseCase) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		vars := mux.Vars(r)

		code := vars["code"]

		err := usecase.Execute(&delete_currency.DeleteCurrencyUseCaseInputDTO{
			Code: code,
		})
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			HandleError(w, err.Error())
			return
		}

		response := struct {
			Message string
		}{
			Message: "Succesfully deleted currency",
		}

		responsebody, err := json.Marshal(response)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		w.Write(responsebody)
	}
}
