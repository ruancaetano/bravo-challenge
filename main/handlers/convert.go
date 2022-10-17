package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ruancaetano/challenge-bravo/domain/usecases/convert_currency"
)

func MakeConvertRequestHandler(r *mux.Router, usecase *convert_currency.ConvertCurrencyUseCase) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		fromCode := r.FormValue("from")
		toCode := r.FormValue("to")
		amount, err := strconv.ParseFloat(r.FormValue("amount"), 64)
		if err != nil {
			HandleError(w, err.Error())
			return
		}

		output, err := usecase.Execute(&convert_currency.ConvertCurrencyUseCaseInputDTO{
			FromCurrencyCode: fromCode,
			ToCurrencyCode:   toCode,
			Amount:           amount,
		})
		if err != nil {
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
