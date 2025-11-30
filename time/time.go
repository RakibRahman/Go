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
	panic("Please implement the IsAfternoonAppointment function")
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	panic("Please implement the Description function")
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	panic("Please implement the AnniversaryDate function")
}

func main() {
	t, err := Schedule("7/25/2019 13:45:00")
	if err != nil {
		// handle the error appropriately
		log.Fatal(err)
	}
	fmt.Println(t)
	fmt.Println(HasPassed("July 25, 2019 13:45:00"))
}
