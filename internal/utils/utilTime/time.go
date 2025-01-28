package utilTime

import (
	"errors"
	"time"
)

func calculateDates(startDate, endDate time.Time, iteration int, days, months, years int) (time.Time, time.Time) {
	return startDate.AddDate(years, months, days*iteration), endDate.AddDate(years, months, days*iteration)
}

func CalculateMonthlyDates(startDate, endDate time.Time, iteration int) (time.Time, time.Time) {
	return calculateDates(startDate, endDate, iteration, 0, 1, 0)
}

func CalculateWeeklyDates(startDate, endDate time.Time, iteration int) (time.Time, time.Time) {
	return calculateDates(startDate, endDate, iteration, 7, 0, 0)
}

func CalculateDailyDates(startDate, endDate time.Time, iteration int) (time.Time, time.Time) {
	return calculateDates(startDate, endDate, iteration, 1, 0, 0)
}

func ValidateDates(startDate, endDate time.Time) error {
	if !endDate.After(startDate) {
		return errors.New("end date must be after start date")
	}
	return nil
}
