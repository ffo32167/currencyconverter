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

	"github.com/ffo32167/currencyconverter/internal/postgres"
)

type Currencyfreaks struct {
	connStr    string
	currencies string
}

type CurrencyfreaksResponse struct {
	Date  string            `json:"date"`
	Base  string            `json:"base"`
	Rates map[string]string `json:"rates"`
}

func New(connStr, currencies string) Currencyfreaks {
	return Currencyfreaks{connStr: connStr, currencies: currencies}
}

func (c Currencyfreaks) Rates() ([]postgres.StorageRate, error) {
	client := http.Client{
		Timeout: 1 * time.Second,
	}
	resp, err := client.Get(c.connStr)
	if err != nil {
		return nil, fmt.Errorf("cant connect with CurrencyFreaks: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("got wrong response status from CurrencyFreaks")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("cant read body of CurrencyFreaks response: %w", err)
	}
	var cfr CurrencyfreaksResponse
	if err := json.Unmarshal([]byte(body), &cfr); err != nil {
		return nil, fmt.Errorf("cant unmarshal data from CurrencyFreaks: %w", err)
	}

	var pgrates []postgres.StorageRate
	date, err := time.Parse("2006-01-02 15:04:05+00", cfr.Date)
	if err != nil {
		return nil, fmt.Errorf("cant parse date from CurrencyFreaks: %w", err)
	}
	pgrates = append(pgrates, postgres.StorageRate{
		RateDate: date,
		CurrCode: cfr.Base,
		Rate:     1,
	})
	for key, val := range cfr.Rates {
		if strings.Contains(c.currencies, key) {
			rate, err := strconv.ParseFloat(val, 64)
			if err != nil {
				return nil, fmt.Errorf("cant parse rate value from CurrencyFreaks: %w", err)
			}
			pgrates = append(pgrates, postgres.StorageRate{RateDate: date, CurrCode: key, Rate: rate})
		}
	}
	return pgrates, nil
}
