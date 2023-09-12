package timeago

import (
	"fmt"
	"testing"
	"time"
)

func TestOfSecondPrecision(t *testing.T) {
	for expected, input := range map[string]time.Time{
		"1 minute from now":   time.Now().Add(63 * time.Second),
		"14 seconds from now": time.Now().Add(14 * time.Second),
		"2 seconds from now":  time.Now().Add(2 * time.Second),
		"1 second from now":   time.Now().Add(1 * time.Second),
		"now":                 time.Now(),
		"1 second ago":        time.Now().Add(-1 * time.Second),
		"2 seconds ago":       time.Now().Add(-2 * time.Second),
		"14 seconds ago":      time.Now().Add(-14 * time.Second),
		"1 minute ago":        time.Now().Add(-63 * time.Second),
	} {
		t.Run(expected, func(t *testing.T) {
			isEqual(t, Of(input, Options{Precision: SecondPrecision}), expected)
		})
	}
}

func TestOfMinutePrecision(t *testing.T) {
	for expected, input := range map[string]time.Time{
		"next hour":           time.Now().Add(63 * time.Minute),
		"14 minutes from now": time.Now().Add(14 * time.Minute),
		"2 minutes from now":  time.Now().Add(2 * time.Minute),
		"1 minute from now":   time.Now().Add(1 * time.Minute),
		"now":                 time.Now(),
		"1 minute ago":        time.Now().Add(-1 * time.Minute),
		"2 minutes ago":       time.Now().Add(-2 * time.Minute),
		"14 minutes ago":      time.Now().Add(-14 * time.Minute),
		"last hour":           time.Now().Add(-63 * time.Minute),
	} {
		t.Run(expected, func(t *testing.T) {
			isEqual(t, Of(input, Options{Precision: MinutePrecision}), expected)
		})
	}
}

func TestOfHourPrecision(t *testing.T) {
	for expected, input := range map[string]time.Time{
		"tomorrow":          time.Now().Add(25 * time.Hour),
		"14 hours from now": time.Now().Add(14 * time.Hour),
		"2 hours from now":  time.Now().Add(2 * time.Hour),
		"next hour":         time.Now().Add(1 * time.Hour),
		"this hour":         time.Now(),
		"last hour":         time.Now().Add(-1 * time.Hour),
		"2 hours ago":       time.Now().Add(-2 * time.Hour),
		"14 hours ago":      time.Now().Add(-14 * time.Hour),
		"yesterday":         time.Now().Add(-25 * time.Hour),
	} {
		t.Run(expected, func(t *testing.T) {
			isEqual(t, Of(input, Options{Precision: HourPrecision}), expected)
		})
	}
}

func TestOfDayPrecision(t *testing.T) {
	for expected, input := range map[string]time.Time{
		"next month":       time.Now().Add(1 * 30 * 24 * time.Hour),
		"14 days from now": time.Now().Add(14 * 24 * time.Hour),
		"2 days from now":  time.Now().Add(38 * time.Hour),
		"tomorrow":         time.Now().Add(25 * time.Hour),
		"today":            time.Now(),
		"yesterday":        time.Now().Add(-25 * time.Hour),
		"2 days ago":       time.Now().Add(-38 * time.Hour),
		"14 days ago":      time.Now().Add(-14 * 24 * time.Hour),
		"last month":       time.Now().Add(-1 * 30 * 24 * time.Hour),
	} {
		t.Run(expected, func(t *testing.T) {
			isEqual(t, Of(input, Options{Precision: DayPrecision}), expected)
		})
	}
}

func TestOfMonthPrecision(t *testing.T) {
	for expected, input := range map[string]time.Time{
		"next year":         time.Now().Add(1 * 365 * 24 * time.Hour),
		"3 months from now": time.Now().Add(3 * 30 * 24 * time.Hour),
		"next month":        time.Now().Add(1 * 30 * 24 * time.Hour),
		"this month":        time.Now(),
		"last month":        time.Now().Add(-1 * 30 * 24 * time.Hour),
		"3 months ago":      time.Now().Add(-3 * 30 * 24 * time.Hour),
		"last year":         time.Now().Add(-1 * 365 * 24 * time.Hour),
	} {
		t.Run(expected, func(t *testing.T) {
			isEqual(t, Of(input, Options{Precision: MonthPrecision}), expected)
		})
	}
}

func TestOfYearPrecision(t *testing.T) {
	for expected, input := range map[string]time.Time{
		"3 years from now": time.Now().Add(3 * 365 * 24 * time.Hour),
		"next year":        time.Now().Add(1 * 365 * 24 * time.Hour),
		"this year":        time.Now(),
		"last year":        time.Now().Add(-1 * 365 * 24 * time.Hour),
		"3 years ago":      time.Now().Add(-3 * 365 * 24 * time.Hour),
	} {
		t.Run(expected, func(t *testing.T) {
			isEqual(t, Of(input, Options{Precision: YearPrecision}), expected)
		})
	}
}

func TestOfCustomFormat(t *testing.T) {
	isEqual(
		t,
		Of(
			time.Now(),
			Options{
				Precision: MinutePrecision,
				Format: Format{
					ThisMinute: "aloha",
				},
			},
		),
		"aloha",
	)
}

func isEqual(tb testing.TB, result, expected string) {
	tb.Helper()
	if expected != result {
		tb.Errorf("expected %q, got %q", expected, result)
	}
}

func ExampleOf() {
	fmt.Println(Of(time.Now().Add(-10 * time.Second)))
	// Output:
	// 10 seconds ago
}
