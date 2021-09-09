package cron

import (
	"fmt"
	"time"
)

type Cron struct {
	date time.Time
	loc  time.Location
	fn   func()
}

func New(date time.Time, loc time.Location, fn func()) Cron {
	return Cron{date: date, loc: loc, fn: fn}
}

func (c Cron) Action(f func()) error {
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

	fmt.Println(firstCallTime)
	if firstCallTime.Before(now) {
		firstCallTime = firstCallTime.Add(time.Hour * 24)
		fmt.Println(firstCallTime)
	}

	duration := firstCallTime.Sub(time.Now().In(&c.loc))

	go func() {
		time.Sleep(duration)
		for {
			f()
			time.Sleep(time.Hour * 24)
		}
	}()

	return nil
}
