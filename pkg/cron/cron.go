package cron

import "time"

type Cron struct {
	hour int
	min  int
	sec  int
	fn   func()
}

func (c Cron) CallFuncAtTime() error {
	time.Sleep(15 * time.Second)
	go func() {
		c.fn()
	}()
	return nil
}
