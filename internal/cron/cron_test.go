package cron

import (
	"context"
	"testing"
	"time"

	"github.com/ffo32167/currencyconverter/internal"
	"go.uber.org/zap"
)

func fn(td time.Duration, s internal.Source, st internal.Storage) error {
	st.Create(context.TODO(), nil)
	return nil
}

type source struct {
	rates []internal.Rate
}

func (s source) Rates() ([]internal.Rate, error) {
	return s.rates, nil
}

type storage struct {
	res chan struct{}
}

func (s storage) Rate(ctx context.Context, date time.Time) ([]internal.Rate, error) {
	return nil, nil
}
func (s storage) Create(ctx context.Context, rates []internal.Rate) error {
	s.res <- struct{}{}
	return nil
}

// результатом Cron является запуск функции с определенным интервалом
func TestAction(t *testing.T) {
	var interval int64
	interval = 1
	countTo := 3

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	date, err := time.Parse("2006-01-02", "2021-10-21")
	if err != nil {
		t.Fatal("cant parse cron date")
	}

	source := source{
		rates: []internal.Rate{
			{RateDate: date, CurrCode: "RUB", Rate: 1},
			{RateDate: date, CurrCode: "USD", Rate: 30.9436},
			{RateDate: date, CurrCode: "EUR", Rate: 26.8343},
			{RateDate: date, CurrCode: "JPY", Rate: 0.231527},
		},
	}
	storage := storage{res: make(chan struct{})}

	loc, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		t.Fatal("cant parse cron location")
	}

	for _, tt := range []Cron{
		{
			date:     date,
			loc:      *loc,
			fn:       fn,
			timeout:  1 * time.Second,
			source:   source,
			storage:  storage,
			log:      logger,
			interval: interval,
		},
	} {
		go tt.Action()

		start := time.Now()
		for i := 0; i < countTo; i++ {
			<-storage.res
		}
		duration := time.Since(start)

		if duration.Seconds() < float64((countTo-1)*int(interval)) || duration.Seconds() > float64(countTo*int(interval)) {
			t.Fatalf("cron call wrong time, duration %s, retry %v", duration, countTo)
		}

		if err != nil {
			t.Fatal("error occured while testing cron: ", err)
		}
	}
}
