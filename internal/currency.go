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
	Rate(ctx context.Context, date string) ([]Rate, error)
	Create(ctx context.Context, rates []Rate) error
}

type Rate struct {
	RateDate time.Time
	CurrCode string
	Rate     float64
}

func Relation(ctx context.Context, storage Storage, date string, curr1, curr2 string) (string, error) {
	data, err := storage.Rate(ctx, date)
	if err != nil {
		return "", fmt.Errorf("cant get rate from storage: %w", err)
	}
	curr1index := sort.Search(len(data), func(i int) bool {
		return data[i].CurrCode >= curr1
	})

	curr2index := sort.Search(len(data), func(i int) bool {
		return data[i].CurrCode >= curr2
	})

	if data[curr2index].Rate == float64(0) {
		return "", errors.New("error, rate of the " + curr2 + " is 0")
	}

	relation := data[curr2index].Rate / data[curr1index].Rate

	return "exchange rate for " + curr1 + " to " + curr2 + " is " +
		strconv.FormatFloat(relation, 'f', 6, 64), nil
}
