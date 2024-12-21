package models

import "time"

type DateRange struct {
	From time.Time
	To   time.Time
}

func (d DateRange) GetFormattedFrom() string {
	return d.From.Format("2006-01-02")
}

func (d DateRange) GetFormattedTo() string {
	return d.To.Format("2006-01-02")
}
