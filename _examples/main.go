package main

import (
	"fmt"
	"time"

	timeago "github.com/caarlos0/timea.go"
)

func main() {
	fmt.Println("# with year precision\n")
	for _, t := range []time.Time{
		time.Now().Add(3 * 365 * 24 * time.Hour),
		time.Now().Add(1 * 365 * 24 * time.Hour),
		time.Now(),
		time.Now().Add(-1 * 365 * 24 * time.Hour),
		time.Now().Add(-3 * 365 * 24 * time.Hour),
	} {
		fmt.Println("* " + timeago.Of(t, timeago.Options{
			Precision: timeago.YearPrecision,
		}))
	}

	fmt.Println("\n\n# with month precision\n")
	for _, t := range []time.Time{
		time.Now().Add(1 * 365 * 24 * time.Hour),
		time.Now().Add(3 * 30 * 24 * time.Hour),
		time.Now().Add(1 * 30 * 24 * time.Hour),
		time.Now(),
		time.Now().Add(-1 * 30 * 24 * time.Hour),
		time.Now().Add(-3 * 30 * 24 * time.Hour),
		time.Now().Add(-1 * 365 * 24 * time.Hour),
	} {
		fmt.Println("* " + timeago.Of(t, timeago.Options{
			Precision: timeago.MonthPrecision,
		}))
	}

	fmt.Println("\n\n# with day precision\n")
	for _, t := range []time.Time{
		time.Now().Add(1 * 30 * 24 * time.Hour),
		time.Now().Add(14 * 24 * time.Hour),
		time.Now().Add(38 * time.Hour),
		time.Now().Add(25 * time.Hour),
		time.Now(),
		time.Now().Add(-25 * time.Hour),
		time.Now().Add(-38 * time.Hour),
		time.Now().Add(-14 * 24 * time.Hour),
		time.Now().Add(-1 * 30 * 24 * time.Hour),
	} {
		fmt.Println("* " + timeago.Of(t, timeago.Options{
			Precision: timeago.DayPrecision,
		}))
	}

	fmt.Println("\n\n# with hour precision\n")
	for _, t := range []time.Time{
		time.Now().Add(25 * time.Hour),
		time.Now().Add(14 * time.Hour),
		time.Now().Add(2 * time.Hour),
		time.Now().Add(1 * time.Hour),
		time.Now(),
		time.Now().Add(-1 * time.Hour),
		time.Now().Add(-2 * time.Hour),
		time.Now().Add(-14 * time.Hour),
		time.Now().Add(-25 * time.Hour),
	} {
		fmt.Println("* " + timeago.Of(t, timeago.Options{
			Precision: timeago.HourPrecision,
		}))
	}

	fmt.Println("\n\n# with minute precision\n")
	for _, t := range []time.Time{
		time.Now().Add(63 * time.Minute),
		time.Now().Add(14 * time.Minute),
		time.Now().Add(2 * time.Minute),
		time.Now().Add(1 * time.Minute),
		time.Now(),
		time.Now().Add(-1 * time.Minute),
		time.Now().Add(-2 * time.Minute),
		time.Now().Add(-14 * time.Minute),
		time.Now().Add(-63 * time.Minute),
	} {
		fmt.Println("* " + timeago.Of(t, timeago.Options{
			Precision: timeago.MinutePrecision,
		}))
	}

	fmt.Println("\n\n# with second precision\n")
	for _, t := range []time.Time{
		time.Now().Add(63 * time.Second),
		time.Now().Add(14 * time.Second),
		time.Now().Add(2 * time.Second),
		time.Now().Add(1 * time.Second),
		time.Now(),
		time.Now().Add(-1 * time.Second),
		time.Now().Add(-2 * time.Second),
		time.Now().Add(-14 * time.Second),
		time.Now().Add(-63 * time.Second),
	} {
		fmt.Println("* " + timeago.Of(t, timeago.Options{
			Precision: timeago.SecondPrecision,
		}))
	}
}
