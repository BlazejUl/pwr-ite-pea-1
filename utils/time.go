package utils

import (
	"time"
)

func GetTimeFromStartInNano(start time.Time) float64 {

	return float64(time.Since(start).Nanoseconds())
}
