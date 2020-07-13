package main

import "time"

type event struct {
	raw []interface{}
}

func newEvent(raw []interface{}) *event {
	return &event{
		raw: raw,
	}
}

func (e *event) Time() (time.Time, error) {
	//"09/Jul/2020 16:34:27",
	return time.Parse("2/Jan/2006 15:04:05", e.date())
}
