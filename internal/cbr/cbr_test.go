package cbr

import (
	"reflect"
	"testing"
	"time"

	"github.com/ffo32167/currencyconverter/internal"
)

func TestRates(t *testing.T) {
	date, err := time.Parse("2006-01-02", "2002-03-02")
	if err != nil {
		t.Fatal("cant parse date: ", err)
	}
	for _, tt := range []struct {
		cbr    Cbr
		result []internal.Rate
	}{
		{cbr: New(
			"https://www.cbr.ru/scripts/XML_daily.asp?date_req=02/03/2002",
			"USD,EUR,RUB,JPY",
			1*time.Second),
			result: []internal.Rate{
				{RateDate: date, CurrCode: "RUB", Rate: 1},
				{RateDate: date, CurrCode: "USD", Rate: 30.9436},
				{RateDate: date, CurrCode: "EUR", Rate: 26.8343},
				{RateDate: date, CurrCode: "JPY", Rate: 0.231527},
			}},
	} {
		res, err := tt.cbr.Rates()
		if err != nil {
			t.Fatal("error occured while cbr.Rates: ", err)
		}

		if !reflect.DeepEqual(res, tt.result) {
			t.Fatalf("expected %v but got %v", tt.result, res)
		}
	}

}
