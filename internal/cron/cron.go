package cron

import (
	"time"

	"github.com/ffo32167/currencyconverter/internal"
	"go.uber.org/zap"
)

type Cron struct {
	date    time.Time
	loc     time.Location
	fn      func(time.Duration, internal.Source, internal.Storage) error
	timeout time.Duration
	source  internal.Source
	storage internal.Storage
	log     *zap.Logger
}

func New(date time.Time,
	loc time.Location,
	fn func(timeout time.Duration, source internal.Source, storage internal.Storage) error,
	timeout time.Duration,
	source internal.Source,
	storage internal.Storage,
	log *zap.Logger) Cron {
	return Cron{date: date, loc: loc, fn: fn, timeout: timeout, source: source, storage: storage, log: log}
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

	c.log.Info("firstCallTime", zap.Time("firstCallTime", firstCallTime))
	c.log.Info("now: 		", zap.Time("now", time.Now().In(&c.loc)))
	c.log.Info("duration: 	", zap.Duration("duration", duration))

	func() {
		time.Sleep(duration)
		for {
			go c.fn(c.timeout, c.source, c.storage)
			c.log.Info("cron call function every ", zap.Int("interval", 24))
			time.Sleep(time.Second * 24)
		}
	}()
	return nil
}
