package currencyfreaks

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/ffo32167/currencyconverter/internal"
)

type Currencyfreaks struct {
	connStr    string
	currencies string
	client     http.Client
}

type CurrencyfreaksResponse struct {
	Date  string            `json:"date"`
	Base  string            `json:"base"`
	Rates map[string]string `json:"rates"`
}

func New(connStr, currencies string, ctxTimeout int64) Currencyfreaks {
	return Currencyfreaks{
		connStr:    connStr,
		currencies: currencies,
		client:     http.Client{Timeout: time.Duration(ctxTimeout) * time.Second}}
}

func (c Currencyfreaks) Rates() ([]internal.Rate, error) {
	resp, err := c.client.Get(c.connStr)
	if err != nil {
		return nil, fmt.Errorf("cant connect with currencyfreaks: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("got wrong response status from currencyfreaks")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("cant read body of currencyfreaks response: %w", err)
	}
	var cfr CurrencyfreaksResponse
	if err := json.Unmarshal([]byte(body), &cfr); err != nil {
		return nil, fmt.Errorf("cant unmarshal data from currencyfreaks: %w", err)
	}

	if len(cfr.Rates) == 0 {
		return nil, errors.New("cant get rates from currencyfreaks")
	}
	return toDomain(cfr, c.currencies)
}

func toDomain(cfr CurrencyfreaksResponse, currencies string) ([]internal.Rate, error) {
	var rates []internal.Rate
	date, err := time.Parse("2006-01-02 15:04:05+00", cfr.Date)
	if err != nil {
		return nil, fmt.Errorf("cant parse date from currencyfreaks: %w", err)
	}
	rates = append(rates, internal.Rate{
		RateDate: date,
		CurrCode: cfr.Base,
		Rate:     1,
	})
	for key, val := range cfr.Rates {
		if strings.Contains(currencies, key) {
			rate, err := strconv.ParseFloat(val, 64)
			if err != nil {
				return nil, fmt.Errorf("cant parse rate value from currencyfreaks: %w", err)
			}
			rates = append(rates, internal.Rate{RateDate: date, CurrCode: key, Rate: rate})
		}
	}
	return rates, nil
}
