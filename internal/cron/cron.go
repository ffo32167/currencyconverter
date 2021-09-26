package cron

import (
	"fmt"
	"time"

	"github.com/ffo32167/currencyconverter/internal"
)

type Cron struct {
	date    time.Time
	loc     time.Location
	fn      func(time.Duration, internal.Source, internal.Storage) error
	timeout time.Duration
	source  internal.Source
	storage internal.Storage
}

func New(date time.Time,
	loc time.Location,
	fn func(timeout time.Duration, source internal.Source, storage internal.Storage) error,
	timeout time.Duration,
	source internal.Source,
	storage internal.Storage) Cron {
	return Cron{date: date, loc: loc, fn: fn, timeout: timeout, source: source, storage: storage}
}

func (c Cron) Action() error {
	now := time.Now().In(&c.loc)
	firstCallTime := time.Date(
		now.Year(),
		now.Month(),
		now.Day(),
		c.date.Hour(),
		c.date.Minute(),
		c.date.Second(),
		0,
		&c.loc,
	)

	duration := firstCallTime.Sub(time.Now().In(&c.loc))

	fmt.Println("firstCallTime: 	", firstCallTime)
	fmt.Println("Now: 		", time.Now().In(&c.loc))
	fmt.Println("duration: 	", duration)

	func() {
		time.Sleep(duration)
		for {
			go c.fn(c.timeout, c.source, c.storage)
			fmt.Println("cron call function")
			time.Sleep(time.Second * 24)
		}
	}()
	return nil
}
