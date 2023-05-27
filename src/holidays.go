package go_nyse_time

import "time"

// The IsWeekend function in Go checks if a given date is a Saturday or Sunday.
func IsWeekend(t time.Time) bool {
	weekday := t.Weekday()
	return weekday == time.Saturday || weekday == time.Sunday
}

// The IsHoliday function in Go checks if a given date is a holiday.
func IsHoliday(t time.Time) bool {
	year := t.Year()
	month := t.Month()
	day := t.Day()

	newYears := newYearsDay(year)
	mlkJr := mlkJrDay(year)
	washingtonsBirthday := washingtonsBirthday(year)
	goodFriday := goodFriday(year)
	memorialDay := memorialDay(year)
	juneteenth := juneteenth(year)
	independenceDay := independenceDay(year)
	laborDay := laborDay(year)
	thanksgivingDay := thanksgivingDay(year)
	christmasDay := christmasDay(year)

	return (fridayNYE(year) && month == time.December && day == 31) ||
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

// The function checks if a given date is an early close day based on specific holidays and weekends.
func IsEarlyClose(t time.Time) bool {
	year := t.Year()
	month := t.Month()
	day := t.Day()

	julyThird := julyThird(year)
	blackFriday := blackFriday(year)
	christmasEveDay := christmasEveDay(year)

	return (month == julyThird.Month() && day == julyThird.Day()) && !IsWeekend(t) ||
		(month == blackFriday.Month() && day == blackFriday.Day()) ||
		(month == christmasEveDay.Month() && day == christmasEveDay.Day() && !IsWeekend(t))
}

// The function returns the nth occurrence of a specific weekday in a given month and year.
func nthWeekday(year int, month time.Month, weekday time.Weekday, n int) time.Time {
	t := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	for i := 0; i < n; i++ {
		t = nextWeekday(t, weekday)
	}
	return t
}

// The function returns the next date that matches a given weekday.
func nextWeekday(t time.Time, weekday time.Weekday) time.Time {
	for t.Weekday() != weekday {
		t = t.AddDate(0, 0, 1)
	}
	return t
}

// The function calculates the date of Easter Sunday for a given year using the Meeus/Jones/Butcher algorithm.
func easterSunday(year int) time.Time {
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

// The function checks if New Year's Eve falls on a Friday in a given year.
func fridayNYE(year int) bool {
	theDay := time.Date(year, time.December, 31, 0, 0, 0, 0, time.UTC).Weekday()
	return theDay == time.Friday
}

// The function returns the date of New Year's Day for a given year, adjusting for cases where it falls on a Sunday.
func newYearsDay(year int) time.Time {
	theDay := time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC).Weekday()

	if theDay == time.Sunday {
		return time.Date(year, time.January, 2, 0, 0, 0, 0, time.UTC)
	}
	return time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)
}

// The function returns the date of the third Monday in January of a given year, which is Martin Luther King Jr. Day in the United States.
func mlkJrDay(year int) time.Time {
	return nthWeekday(year, time.January, time.Monday, 3)
}

// The function returns the date of the third Monday in February for a given year.
func washingtonsBirthday(year int) time.Time {
	return nthWeekday(year, time.February, time.Monday, 3)
}

// The function returns the date of Good Friday for a given year.
func goodFriday(year int) time.Time {
	return easterSunday(year).AddDate(0, 0, -2)
}

// The function returns the date of the last Monday in May for a given year.
func memorialDay(year int) time.Time {
	return nthWeekday(year, time.May, time.Monday, -1)
}

// The function returns the date of Juneteenth for a given year in Go.
func juneteenth(year int) time.Time {
	return time.Date(year, time.June, 19, 0, 0, 0, 0, time.UTC)
}

// The function returns a time object representing July 3rd of a given year.
func julyThird(year int) time.Time {
	return time.Date(year, time.July, 3, 0, 0, 0, 0, time.UTC)
}

// The function returns the date of Independence Day for a given year, adjusting for weekends.
func independenceDay(year int) time.Time {
	theDay := time.Date(year, time.July, 4, 0, 0, 0, 0, time.UTC).Weekday()
	if theDay == time.Saturday {
		return time.Date(year, time.July, 3, 0, 0, 0, 0, time.UTC)
	}
	if theDay == time.Sunday {
		return time.Date(year, time.July, 5, 0, 0, 0, 0, time.UTC)
	}
	return time.Date(year, time.July, 4, 0, 0, 0, 0, time.UTC)
}

// The function "laborDay" returns the date of the first Monday in September of a given year.
func laborDay(year int) time.Time {
	return nthWeekday(year, time.September, time.Monday, 1)
}

// The function returns the date of the fourth Thursday in November for a given year, representing Thanksgiving Day in the US.
func thanksgivingDay(year int) time.Time {
	return nthWeekday(year, time.November, time.Thursday, 4)
}

// The function "blackFriday" returns the date of the fourth Friday in November of a given year.
func blackFriday(year int) time.Time {
	return nthWeekday(year, time.November, time.Friday, 4)
}

// The function returns a time object representing Christmas Eve day for a given year.
func christmasEveDay(year int) time.Time {
	return time.Date(year, time.December, 24, 0, 0, 0, 0, time.UTC)
}

// The function returns the date of Christmas Day for a given year, adjusting for weekends.
func christmasDay(year int) time.Time {
	theDay := time.Date(year, time.December, 25, 0, 0, 0, 0, time.UTC).Weekday()
	if theDay == time.Saturday {
		return time.Date(year, time.December, 24, 0, 0, 0, 0, time.UTC)
	}
	if theDay == time.Sunday {
		return time.Date(year, time.December, 26, 0, 0, 0, 0, time.UTC)
	}
	return time.Date(year, time.December, 25, 0, 0, 0, 0, time.UTC)
}
