package nyse_time

import "time"

/**SECTION - Exchange Trading Hours
 * @desc The following functions are used to determine if the market is open.
 */

/**ANCHOR - IsNormalHours
 * @param t time.Time
 * @return bool
 * @desc Returns true if the market is open on the given date, false otherwise.
 * @note The market is open from 9:30am to 4:00pm EST/EDT on weekdays.
 * @note The market is closed on weekends and holidays.
 * @note The market closes at 1:00pm EST/EDT on the day before Thanksgiving.
 * @note The market closes at 1:00pm EST/EDT on Christmas Eve.
 * @note The market closes at 1:00pm EST/EDT on the day before Independence Day.
 */
func IsNormalHours(t time.Time) bool {
	//NOTE - The market always opens at 9:30AM EST except weekends and holidays.
	openTime := time.Date(t.Year(), t.Month(), t.Day(), 14, 30, 0, 0, time.UTC)
	//NOTE - The market closes at 1PM EST on certain days.
	earlyCloseTime := time.Date(t.Year(), t.Month(), t.Day(), 18, 00, 0, 0, time.UTC)
	//NOTE - The market closes at 4PM EST on normal days.
	closeTime := time.Date(t.Year(), t.Month(), t.Day(), 21, 00, 0, 0, time.UTC)

	if IsEarlyClose(t) {
		return t.After(openTime) && t.Before(earlyCloseTime)
	}

	return t.After(openTime) && t.Before(closeTime)
}
