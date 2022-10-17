package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/ruancaetano/challenge-bravo/domain/services"
	"github.com/ruancaetano/challenge-bravo/infra/cache"
)

const API_URL = "https://economia.awesomeapi.com.br/last"

type AwesomeCurrencyConverterService struct {
	cache cache.CacheInterface
}

func NewAwesomeCurrencyConverterService(cache cache.CacheInterface) *AwesomeCurrencyConverterService {
	return &AwesomeCurrencyConverterService{
		cache,
	}
}

func (s *AwesomeCurrencyConverterService) Convert(fromCode string, toCode string, amount float64) (*services.CurrencyConverterServiceResponse, error) {
	responseBody := s.cache.Get(fromCode + toCode)

	if responseBody == nil {
		result, err := s.getConversionFromApi(fromCode, toCode)
		if err != nil {
			return nil, err
		}

		responseBody = result
		s.cache.Set(fromCode+toCode, responseBody, time.Minute)
	}

	mappedResponseBody := responseBody.(map[string]map[string]string)

	bid, err := strconv.ParseFloat(mappedResponseBody[fromCode+toCode]["bid"], 64)
	if err != nil {
		return nil, err
	}

	return &services.CurrencyConverterServiceResponse{
		FromCurrencyCode: fromCode,
		ToCurrencyCode:   toCode,
		Value:            amount * bid,
	}, nil
}

func (s *AwesomeCurrencyConverterService) getConversionFromApi(fromCode string, toCode string) (interface{}, error) {
	resp, err := http.Get(fmt.Sprintf("%s/%s-%s", API_URL, fromCode, toCode))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("unsupported conversion")
	}

	reponseBody := make(map[string]map[string]string)

	err = json.NewDecoder(resp.Body).Decode(&reponseBody)

	if err != nil {
		return nil, err
	}

	return reponseBody, nil
}
