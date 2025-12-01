package main

import (
	"fmt"
	"log"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) (time.Time, error) {
	layout := "1/2/2006 15:04:05"
	return time.Parse(layout, date)
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	t, err := time.Parse("January 2, 2006 15:04:05", date)
	if err != nil {
		return false
	}
	return time.Now().After(t)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	t, err := time.Parse(layout, date)
	if err != nil {
		return false
	}

	if t.Hour() >= 12 && t.Hour() < 18 {
		return true
	}
	return false
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "1/2/2006 15:04:05"
	t, err := time.Parse(layout, date)
	if err != nil {
		return "Invalid date"
	}
	return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%02d.",
		t.Weekday(), t.Month(), t.Day(), t.Year(), t.Hour(), t.Minute())
}

func DescriptionV2(date string) string {
	layout := "1/2/2006 15:04:05"
	t, err := time.Parse(layout, date)
	if err != nil {
		return "Invalid date"
	}
	return fmt.Sprintf("You have an appointment on %s.",
		t.Format("Monday, January 2, 2006, at 15:04"))
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	year := time.Now().Year()
	return time.Date(year, time.September, 15, 0, 0, 0, 0, time.UTC)
}

func main() {
	t, err := Schedule("7/25/2019 13:45:00")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(t)
	fmt.Println(HasPassed("July 25, 2019 13:45:00"))
	isAfterNoon := IsAfternoonAppointment("Thursday, July 25, 2019 13:45:00")
	fmt.Println(isAfterNoon)
	des := Description("7/25/2019 13:45:00")
	des2 := DescriptionV2("7/12/2025 3:45:00")
	fmt.Println(des)
	fmt.Println(des2)
	fmt.Println(AnniversaryDate())
}
