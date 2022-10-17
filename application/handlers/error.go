package handlers

import (
	"encoding/json"
	"net/http"
)

func HandleError(w http.ResponseWriter, message string) {
	errorMessage := struct {
		Message string `json:"message"`
	}{
		Message: message,
	}

	jsondata, err := json.Marshal(errorMessage)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Write(jsondata)
}
