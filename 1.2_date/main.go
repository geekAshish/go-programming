package main

import (
	"fmt"
	"time"
)

func main() {
	// date, time, duration
	// time.Time, time.Duration
	// time.Date(year int, month Month, day, hour, min, sec int, nsec int, loc *Location) Time
	// time.Now() -> current date and time
	// time.Unix(sec int64, nsec int64) Time
	// time.Sleep(d Duration)
	// time.Parse(layout string, value string) (Time, error)
	// time.UTC() -> UTC timezone
	// time.Local() -> Local timezone
	// time.LoadLocation(name string) (*Location, error)

	presentDate := time.Now()
	fmt.Println("Present date:", presentDate)
	fmt.Println("Present date:", presentDate.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC)
	fmt.Println("Created date:", createdDate)
	fmt.Println("Created date:", createdDate.Format("01-02-2006 15:04:05 Monday"))
}