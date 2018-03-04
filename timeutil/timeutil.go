package timeutil

import (
	"strings"
	"time"
)

func ParseStringTime(str string) time.Time {
	var ret time.Time
	if strings.ContainsRune(str, ':') {
		ret, _ = time.Parse("2006.01.02 03:04", str)
		return ret
	}
	ret, _ = time.Parse("2006.01.02", str)
	return ret
}

func CalcTimeDura(time1, time2 time.Time) float64 {
	dura := time2.Sub(time1)
	return dura.Minutes()
}
