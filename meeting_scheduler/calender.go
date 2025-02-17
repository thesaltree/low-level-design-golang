package main

import (
	"errors"
	"time"
)

type interval struct {
	id     int
	date   time.Time
	start  time.Time
	end    time.Time
	status BookingStatus
}

type Calendar struct {
	interval map[int]*interval
}

func (c *Calendar) isFree(dur *interval) bool {
	// Check if the calendar is free for the given interval, you can add logic to better check if the interval is free
	for _, ci := range c.interval {
		// check if time is colliding
		if ci.status == BOOKED && ci.start.Before(dur.end) && ci.end.After(dur.start) {
			return false
		}
	}
	return true
}

func (c *Calendar) setIntervalStatus(durId int, status BookingStatus) {
	c.interval[durId].status = status
}

func (c *Calendar) bookInterval(dur *interval) error {
	if c.isFree(dur) {
		c.interval[dur.id] = dur
		c.setIntervalStatus(dur.id, BOOKED)
		return nil
	}

	return errors.New("Interval is not free")
}

func (c *Calendar) cancelInterval(durId int) error {
	if val, ok := c.interval[durId]; ok {
		c.setIntervalStatus(val.id, CANCEL)
		return nil
	}

	return errors.New("Interval is not free")
}
