package nyse

import "time"

/**SECTION - Holiday Functions
 * @desc The following functions are used by the holiday algorithms.
 */

/**ANCHOR - IsWeekend
 * @param t time.Time
 * @return bool
 * @desc Returns true if the time is a weekend, false otherwise.
 */
func IsWeekend(t time.Time) bool {
	weekday := t.Weekday()
	return weekday == time.Saturday || weekday == time.Sunday
}

/**ANCHOR - IsHoliday
 * @param t time.Time
 * @return bool
 * @desc Returns true if the date is a holiday, false otherwise.
 */
func IsHoliday(t time.Time) bool {
	year := t.Year()
	month := t.Month()
	day := t.Day()

	newYears := NewYearsDay(year)
	mlkJr := MlkJrDay(year)
	washingtonsBirthday := WashingtonsBirthday(year)
	goodFriday := GoodFriday(year)
	memorialDay := MemorialDay(year)
	juneteenth := Juneteenth(year)
	independenceDay := IndependenceDay(year)
	laborDay := LaborDay(year)
	thanksgivingDay := ThanksgivingDay(year)
	christmasDay := ChristmasDay(year)

	return (FridayNYE(year) && month == time.December && day == 31) ||
		(month == newYears.Month() && day == newYears.Day()) ||
		(month == mlkJr.Month() && day == mlkJr.Day()) ||
		(month == washingtonsBirthday.Month() && day == washingtonsBirthday.Day()) ||
		(month == goodFriday.Month() && day == goodFriday.Day()) ||
		(month == memorialDay.Month() && day == memorialDay.Day()) ||
		(month == juneteenth.Month() && day == juneteenth.Day()) ||
		(month == independenceDay.Month() && day == independenceDay.Day()) ||
		(month == laborDay.Month() && day == laborDay.Day()) ||
		(month == thanksgivingDay.Month() && day == thanksgivingDay.Day()) ||
		(month == christmasDay.Month() && day == christmasDay.Day())
}

/**ANCHOR - IsEarlyClose
 * @param t time.Time
 * @return bool
 * @desc Returns true if the day is an early close day, false otherwise.
 * @note The market closes at 1:00pm EST/EDT on the day before Thanksgiving.
 * @note The market closes at 1:00pm EST/EDT on Christmas Eve.
 * @note The market closes at 1:00pm EST/EDT on the day before Independence Day.
 */
func IsEarlyClose(t time.Time) bool {
	year := t.Year()
	month := t.Month()
	day := t.Day()

	julyThird := JulyThird(year)
	blackFriday := BlackFriday(year)
	christmasEveDay := ChristmasEveDay(year)

	return (month == julyThird.Month() && day == julyThird.Day()) && !IsWeekend(t) ||
		(month == blackFriday.Month() && day == blackFriday.Day()) ||
		(month == christmasEveDay.Month() && day == christmasEveDay.Day() && !IsWeekend(t))
}

/**ANCHOR - nthWeekday
 * @param year int
 * @param month time.Month
 * @param weekday time.Weekday
 * @param n int
 * @return time.Time
 * @desc Returns the nth weekday of the month.
 */
func nthWeekday(year int, month time.Month, weekday time.Weekday, n int) time.Time {
	t := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	for i := 0; i < n; i++ {
		t = nextWeekday(t, weekday)
	}
	return t
}

/**ANCHOR - nextWeekday
 * @param t time.Time
 * @param weekday time.Weekday
 * @return time.Time
 * @desc Returns the next weekday after t.
 */
func nextWeekday(t time.Time, weekday time.Weekday) time.Time {
	for t.Weekday() != weekday {
		t = t.AddDate(0, 0, 1)
	}
	return t
}

/**ANCHOR - EasterSunday
 * @param year int
 * @return time.Time
 * @desc Returns the date of Easter Sunday.
 * @note This holiday is also known as Easter Day.
 * @note This algorithm is based on the algorithm found at
 * LINK - http://www.tondering.dk/claus/cal/easter.php
 */
func EasterSunday(year int) time.Time {
	a := year % 19
	b := year / 100
	c := year % 100
	d := b / 4
	e := b % 4
	f := (b + 8) / 25
	g := (b - f + 1) / 3
	h := (19*a + b - d - g + 15) % 30
	i := c / 4
	k := c % 4
	l := (32 + 2*e + 2*i - h - k) % 7
	m := (a + 11*h + 22*l) / 451
	month := (h + l - 7*m + 114) / 31
	day := ((h + l - 7*m + 114) % 31) + 1
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

//!SECTION

/**SECTION - Holiday Algorithms
 * @desc The following functions are used to calculate holidays.
 */

/**ANCHOR - FridayNYE
 * @param year int
 * @return bool
 * @desc Returns true if NYE falls on a Friday.
 * @note This function is only used if the day before New Year's Day is a Friday.
 */
func FridayNYE(year int) bool {
	theDay := time.Date(year, time.December, 31, 0, 0, 0, 0, time.UTC).Weekday()
	return theDay == time.Friday
}

/**ANCHOR - NewYearsDay
 * @param year int
 * @return time.Time
 * @desc Returns the observation date of New Year's Day.
 * @note This holiday is also known as New Years.
 */
func NewYearsDay(year int) time.Time {
	theDay := time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC).Weekday()

	if theDay == time.Sunday {
		return time.Date(year, time.January, 2, 0, 0, 0, 0, time.UTC)
	}
	return time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)
}

/**ANCHOR - MlkJrDay
 * @param year int
 * @return time.Time
 * @desc Returns the date of Martin Luther King Jr. Day.
 * @note This holiday is also known as MLK Jr. Day.
 */
func MlkJrDay(year int) time.Time {
	return nthWeekday(year, time.January, time.Monday, 3)
}

/**ANCHOR - WashingtonsBirthday
 * @param year int
 * @return time.Time
 * @desc Returns the date of Washington's Birthday.
 * @note This holiday is also known as Presidents' Day.
 */
func WashingtonsBirthday(year int) time.Time {
	return nthWeekday(year, time.February, time.Monday, 3)
}

/**ANCHOR - GoodFriday
 * @param year int
 * @return time.Time
 * @desc Returns the date of Good Friday.
 * @note This holiday is also known as Easter Friday.
 * @note This holiday is also known as Holy Friday.
 * @note This holiday is also known as Great Friday.
 * @note This holiday is also known as Easter Eve.
 */
func GoodFriday(year int) time.Time {
	return EasterSunday(year).AddDate(0, 0, -2)
}

/**ANCHOR - MemorialDay
 * @param year int
 * @return time.Time
 * @desc Returns the date of Memorial Day.
 * @note This holiday is also known as Decoration Day.
 * @note This holiday is also known as May Day.
 */
func MemorialDay(year int) time.Time {
	return nthWeekday(year, time.May, time.Monday, -1)
}

/**ANCHOR - Juneteenth
 * @param year int
 * @return time.Time
 * @desc Returns the date of Juneteenth.
 * @note This holiday is also known as Emancipation Day.
 * @note This holiday is also known as Freedom Day.
 * @note This holiday is also known as Jubilee Day.
 * @note This holiday is also known as Liberation Day.
 */
func Juneteenth(year int) time.Time {
	return time.Date(year, time.June, 19, 0, 0, 0, 0, time.UTC)
}

/**ANCHOR - JulyThird
 * @param year int
 * @return time.Time
 * @desc Returns the day before Independence Day.
 * @note If the day before Independence Day is a weekday, then the exchange closes early that day.
 */
func JulyThird(year int) time.Time {
	return time.Date(year, time.July, 3, 0, 0, 0, 0, time.UTC)
}

/**ANCHOR - IndependenceDay
 * @param year int
 * @return time.Time
 * @desc Returns the observation date of Independence Day.
 * @note This holiday is also known as Fourth of July.
 */
func IndependenceDay(year int) time.Time {
	theDay := time.Date(year, time.July, 4, 0, 0, 0, 0, time.UTC).Weekday()
	if theDay == time.Saturday {
		return time.Date(year, time.July, 3, 0, 0, 0, 0, time.UTC)
	}
	if theDay == time.Sunday {
		return time.Date(year, time.July, 5, 0, 0, 0, 0, time.UTC)
	}
	return time.Date(year, time.July, 4, 0, 0, 0, 0, time.UTC)
}

/**ANCHOR - LaborDay
 * @param year int
 * @return time.Time
 * @desc Returns the date of Labor Day.
 */
func LaborDay(year int) time.Time {
	return nthWeekday(year, time.September, time.Monday, 1)
}

/**ANCHOR - ThanksgivingDay
 * @param year int
 * @return time.Time
 * @desc Returns the date of Thanksgiving Day.
 */
func ThanksgivingDay(year int) time.Time {
	return nthWeekday(year, time.November, time.Thursday, 4)
}

/**ANCHOR - BlackFriday
 * @param year int
 * @return time.Time
 * @desc Returns the date of the Friday after Thanksgiving Day.
 * @note This holiday is also known as Black Friday.
 * @note The exchange closes early on this day.
 */
func BlackFriday(year int) time.Time {
	return nthWeekday(year, time.November, time.Friday, 4)
}

/**ANCHOR - ChristmasEveDay
 * @param year int
 * @return time.Time
 * @desc Returns the date of Christmas Eve.
 * @note If this is a weekday the exchange closes early on this day.
 */
func ChristmasEveDay(year int) time.Time {
	return time.Date(year, time.December, 24, 0, 0, 0, 0, time.UTC)
}

/**ANCHOR - ChristmasDay
 * @param year int
 * @return time.Time
 * @desc Returns the observation date of Christmas Day.
 * @note This holiday is also known as Christmas.
 * @note This holiday is also known as Xmas.
 */
func ChristmasDay(year int) time.Time {
	theDay := time.Date(year, time.December, 25, 0, 0, 0, 0, time.UTC).Weekday()
	if theDay == time.Saturday {
		return time.Date(year, time.December, 24, 0, 0, 0, 0, time.UTC)
	}
	if theDay == time.Sunday {
		return time.Date(year, time.December, 26, 0, 0, 0, 0, time.UTC)
	}
	return time.Date(year, time.December, 25, 0, 0, 0, 0, time.UTC)
}

//!SECTION
