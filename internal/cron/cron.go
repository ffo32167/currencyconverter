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
			go c.fn()
			time.Sleep(time.Second * 4)
		}
	}()
	return nil
}
