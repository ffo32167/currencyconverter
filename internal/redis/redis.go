package redis

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/ffo32167/currencyconverter/internal"
	"github.com/go-redis/redis/v8"
)

type Redis struct {
	client *redis.Client
}
type RedisRate struct {
	RateDate time.Time `json:"date"`
	CurrCode string    `json:"curr_code"`
	Rate     float64   `json:"rate"`
}

func New(args []string) (Redis, error) {
	db, err := strconv.Atoi(args[2])
	if err != nil {
		return Redis{}, fmt.Errorf("cant parse redis db parametr: %w", err)
	}
	return Redis{client: redis.NewClient(
			&redis.Options{
				Addr:     args[0],
				Password: args[1],
				DB:       db})},
		nil
}

func (r Redis) Create(ctx context.Context, internalRates []internal.Rate) error {
	for _, v := range internalRates {
		err := r.client.Set(ctx, encode(v), v.Rate, 0).Err()
		if err != nil {
			return fmt.Errorf("cant parse redis db parametr: %w", err)
		}
	}
	return nil
}

func (r Redis) Rate(ctx context.Context, date time.Time) ([]internal.Rate, error) {
	var result []internal.Rate
	iter := r.client.Scan(ctx, 0, date.Format("2006-01-02")+"*", 0).Iterator()
	for iter.Next(ctx) {
		val, err := r.client.Get(ctx, iter.Val()).Result()
		if err != nil {
			return nil, fmt.Errorf("cant get data from redis: %w", err)
		}
		date, curr, err := decode(iter.Val())
		if err != nil {
			return nil, fmt.Errorf("cant decode date,currency from redis: %w", err)
		}
		rate, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return nil, fmt.Errorf("cant decode rate from redis: %w", err)
		}
		result = append(result, internal.Rate{RateDate: date, CurrCode: curr, Rate: rate})
	}
	err := iter.Err()
	if err != nil {
		return nil, fmt.Errorf("error while scan redis: %w", err)
	}
	return result, nil
}

func encode(v internal.Rate) string {
	return v.RateDate.Format("2006-01-02") + "_" + v.CurrCode
}

func decode(s string) (time.Time, string, error) {
	slice := strings.Split(s, "_")
	date, err := time.Parse("2006-01-02", slice[0])
	if err != nil {
		return date, "", fmt.Errorf("error while parse redis result: %s error: %w", s, err)
	}
	return date, slice[1], nil
}
