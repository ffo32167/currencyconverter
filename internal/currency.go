package internal

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strconv"
	"time"
)

type Storage interface {
	Rate(ctx context.Context, date time.Time) ([]Rate, error)
	Create(ctx context.Context, rates []Rate) error
}

type Source interface {
	Rates() ([]Rate, error)
}

type Rate struct {
	RateDate time.Time
	CurrCode string
	Rate     float64
}

func Sync(timeout int64, source Source, storage Storage) error {
	ctx := context.Background()
	rates, err := source.Rates()
	if err != nil {
		return err
	}
	err = storage.Create(ctx, rates)
	if err != nil {
		return err
	}
	return nil
}

func Relation(ctx context.Context, storage Storage, date time.Time, curr1, curr2 string) (string, error) {
	rates, err := storage.Rate(ctx, date)
	if err != nil {
		return "", fmt.Errorf("cant get rate from storage: %w", err)
	}
	curr1index := sort.Search(len(rates), func(i int) bool {
		return rates[i].CurrCode >= curr1
	})

	curr2index := sort.Search(len(rates), func(i int) bool {
		return rates[i].CurrCode >= curr2
	})

	if rates[curr2index].Rate == float64(0) {
		return "", errors.New("error, rate of the " + curr2 + " is 0")
	}

	relation := rates[curr2index].Rate / rates[curr1index].Rate

	return "exchange rate for " + curr1 + " to " + curr2 + " is " +
		strconv.FormatFloat(relation, 'f', 6, 64), nil
}

func Rates(ctx context.Context, storage Storage, date time.Time) ([]Rate, error) {
	return storage.Rate(ctx, date)
}
