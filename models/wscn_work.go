package models

import "time"

type WscnWork struct {
	ID         int64 `gorm:"primary"`
	Date       time.Time
	GetUpTime  time.Time
	GetOffRoom time.Time
	Busline    string
	BusGo      time.Time
	MetroArr   time.Time
	OfficeArr  time.Time
	Duration   float64
	Mark       string
}
