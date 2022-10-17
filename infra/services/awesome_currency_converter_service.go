package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ruancaetano/challenge-bravo/domain/services"
)

const API_URL = "https://economia.awesomeapi.com.br/last"

type AwesomeCurrencyConverterService struct{}

func NewAwesomeCurrencyConverterService() *AwesomeCurrencyConverterService {
	return &AwesomeCurrencyConverterService{}
}

func (s *AwesomeCurrencyConverterService) Convert(fromCode string, toCode string, amount float64) (*services.CurrencyConverterServiceResponse, error) {
	resp, err := http.Get(fmt.Sprintf("%s/%s-%s", API_URL, fromCode, toCode))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Unsupported conversion")
	}

	reponseBody := make(map[string]map[string]string)

	err = json.NewDecoder(resp.Body).Decode(&reponseBody)

	if err != nil {
		return nil, err
	}

	bid, err := strconv.ParseFloat(reponseBody[fromCode+toCode]["bid"], 64)
	if err != nil {
		return nil, err
	}

	return &services.CurrencyConverterServiceResponse{
		FromCurrencyCode: fromCode,
		ToCurrencyCode:   toCode,
		Value:            amount * bid,
	}, nil
}
