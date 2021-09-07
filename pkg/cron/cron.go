package cron

import "time"

type Cron struct {
	date time.Time
	fn   func()
}

func (c Cron) CallFuncAtTime() error {
	time.Sleep(15 * time.Second)
	go func() {
		c.fn()
	}()
	return nil
}
