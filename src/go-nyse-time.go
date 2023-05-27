package go_nyse_time

import (
	"time"
)

// The function checks if the market is currently open based on the current time, excluding weekends and holidays and checking if it's within normal trading hours.
func IsMarketOpen() bool {
	currentTime := time.Now().UTC()
	return !IsWeekend(currentTime) && !IsHoliday(currentTime) && IsNormalHours(currentTime)
}

// The function checks if the market is open during normal hours on a given date, excluding weekends and holidays.
func IsMarketOpenCustom(t time.Time) bool {
	return !IsWeekend(t) && !IsHoliday(t) && IsNormalHours(t)
}
