package main

import (
	"errors"
	"time"
)

type interval struct {
	date  time.Time
	start time.Time
	end   time.Time
}

type Calendar struct {
	interval []interval
}

func (c *Calendar) isFree(dur interval) bool {
	// Check if the calendar is free for the given interval, you can add logic to better check if the interval is free
	for _, ci := range c.interval {
		if ci.date.Equal(dur.date) && ci.start.Equal(dur.start) && ci.end.Equal(dur.end) {
			return false
		}
	}
	return true
}

func (c *Calendar) bookInterval(dur interval) error {
	if c.isFree(dur) {
		c.interval = append(c.interval, dur)
		return nil
	}

	return errors.New("Interval is not free")
}
