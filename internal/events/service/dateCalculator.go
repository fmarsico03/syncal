package service

import "time"

func calculateDates(startDate, endDate time.Time, iteration int, days, months, years int) (time.Time, time.Time) {
	return startDate.AddDate(years, months, days*iteration), endDate.AddDate(years, months, days*iteration)
}

func calculateMonthlyDates(startDate, endDate time.Time, iteration int) (time.Time, time.Time) {
	return calculateDates(startDate, endDate, iteration, 0, 1, 0)
}

func calculateWeeklyDates(startDate, endDate time.Time, iteration int) (time.Time, time.Time) {
	return calculateDates(startDate, endDate, iteration, 7, 0, 0)
}

func calculateDailyDates(startDate, endDate time.Time, iteration int) (time.Time, time.Time) {
	return calculateDates(startDate, endDate, iteration, 1, 0, 0)
}
