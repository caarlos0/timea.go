// Package timeago provides a simple library to format a time in a "time ago" manner.
package timeago

import (
	"fmt"
	"reflect"
	"time"
)

// Precision define the minimun amount of time to be considered.
type Precision uint

const (
	// SecondPrecision is the second precision.
	SecondPrecision Precision = iota

	// MinutePrecision is the minute precision.
	MinutePrecision

	// HourPrecision is the hour precision.
	HourPrecision

	// DayPrecision is the day precision.
	DayPrecision

	// MonthPrecision is the month precision.
	MonthPrecision

	// YearPrecision is the year precision.
	YearPrecision
)

// Options define the options of the library.
type Options struct {
	Precision Precision
	Format    Format
}

// Of returns the string representation of the given time with the given options.
func Of(t time.Time, options ...Options) string {
	opt := Options{
		Precision: SecondPrecision,
		Format:    DefaultFormat,
	}

	for _, o := range options {
		if o.Precision != 0 {
			opt.Precision = o.Precision
		}
		if !reflect.DeepEqual(o.Format, Format{}) {
			opt.Format = o.Format
		}
	}

	switch opt.Precision {
	case SecondPrecision:
		seconds := time.Since(t).Round(time.Second).Seconds()
		if seconds == 0 {
			return opt.Format.ThisSecond
		}
		if seconds == 1 {
			return opt.Format.LastSecond
		}
		if seconds < 60 {
			return fmt.Sprintf(opt.Format.SecondsAgo, int(seconds))
		}
		return Of(t, Options{
			Precision: MinutePrecision,
		})
	case MinutePrecision:
		minutes := time.Since(t).Round(time.Minute).Minutes()
		if minutes == 0 {
			return opt.Format.ThisMinute
		}
		if minutes == 1 {
			return opt.Format.LastMinute
		}
		if minutes < 60 {
			return fmt.Sprintf(opt.Format.MinutesAgo, int(minutes))
		}
		return Of(t, Options{
			Precision: HourPrecision,
		})
	case HourPrecision:
		hours := time.Since(t).Round(time.Hour).Hours()
		if hours == 0 {
			return opt.Format.ThisHour
		}
		if hours == 1 {
			return opt.Format.LastHour
		}
		if hours < 24 {
			return fmt.Sprintf(opt.Format.HoursAgo, int(hours))
		}
		return Of(t, Options{
			Precision: DayPrecision,
		})
	case DayPrecision:
		days := time.Since(t).Round(time.Hour*24).Hours() / 24
		if days == 0 {
			return opt.Format.Today
		}
		if days == 1 {
			return opt.Format.Yesterday
		}
		if days < 30 {
			return fmt.Sprintf(opt.Format.DaysAgo, int(days))
		}
		return Of(t, Options{
			Precision: MonthPrecision,
		})
	case MonthPrecision:
		months := time.Since(t).Round(time.Hour*24*30).Hours() / (24 * 30)
		if months == 0 {
			return opt.Format.ThisMonth
		}
		if months == 1 {
			return opt.Format.LastMonth
		}
		if months < 12 {
			return fmt.Sprintf(opt.Format.MonthsAgo, int(months))
		}
		return Of(t, Options{
			Precision: YearPrecision,
		})
	case YearPrecision:
		years := time.Since(t).Round(time.Hour*24*365).Hours() / (24 * 365)
		if years == 0 {
			return opt.Format.ThisYear
		}
		if years == 1 {
			return opt.Format.LastYear
		}
		return fmt.Sprintf(opt.Format.YearsAgo, int(years))
	}

	// this should never happen
	return t.String()
}

// Format is the format of the string returned by the library.
type Format struct {
	ThisSecond string
	LastSecond string
	SecondsAgo string

	ThisMinute string
	LastMinute string
	MinutesAgo string

	ThisHour string
	LastHour string
	HoursAgo string

	Today     string
	Yesterday string
	DaysAgo   string

	ThisMonth string
	LastMonth string
	MonthsAgo string

	ThisYear string
	LastYear string
	YearsAgo string
}

// DefaultFormat is the default format of the string returned by the library.
var DefaultFormat = Format{
	ThisSecond: "now",
	LastSecond: "1 second ago",
	SecondsAgo: "%d seconds ago",

	ThisMinute: "now",
	LastMinute: "1 minute ago",
	MinutesAgo: "%d minutes ago",

	ThisHour: "this hour",
	LastHour: "last hour",
	HoursAgo: "%d hours ago",

	Today:     "today",
	Yesterday: "yesterday",
	DaysAgo:   "%d days ago",

	ThisMonth: "this month",
	LastMonth: "last month",
	MonthsAgo: "%d months ago",

	ThisYear: "this year",
	LastYear: "last year",
	YearsAgo: "%d years ago",
}
