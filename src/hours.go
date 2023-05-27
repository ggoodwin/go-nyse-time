package go_nyse_time

import "time"

// The function checks if a given time falls within normal trading hours of the stock market.
func IsNormalHours(t time.Time) bool {
	openTime := time.Date(t.Year(), t.Month(), t.Day(), 14, 30, 0, 0, time.UTC)
	earlyCloseTime := time.Date(t.Year(), t.Month(), t.Day(), 18, 00, 0, 0, time.UTC)
	closeTime := time.Date(t.Year(), t.Month(), t.Day(), 21, 00, 0, 0, time.UTC)

	if IsEarlyClose(t) {
		return t.After(openTime) && t.Before(earlyCloseTime)
	}

	return t.After(openTime) && t.Before(closeTime)
}
