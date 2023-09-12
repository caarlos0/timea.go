// Package timeago provides a simple library to format a time in a "time ago" manner.
package timeago

import (
	"fmt"
	"reflect"
	"time"
)

// Precision define the minimum amount of time to be considered.
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
		if seconds > -60 && seconds < -1 {
			return fmt.Sprintf(opt.Format.SecondsFromNow, positive(seconds))
		}
		if seconds == -1 {
			return opt.Format.NextSecond
		}
		if seconds == 0 {
			return opt.Format.ThisSecond
		}
		if seconds == 1 {
			return opt.Format.LastSecond
		}
		if seconds > 0 && seconds < 60 {
			return fmt.Sprintf(opt.Format.SecondsAgo, positive(seconds))
		}
		return Of(t, Options{
			Precision: MinutePrecision,
		})
	case MinutePrecision:
		minutes := time.Since(t).Round(time.Minute).Minutes()
		if minutes > -60 && minutes < -1 {
			return fmt.Sprintf(opt.Format.MinutesFromNow, positive(minutes))
		}
		if minutes == -1 {
			return opt.Format.NextMinute
		}
		if minutes == 0 {
			return opt.Format.ThisMinute
		}
		if minutes == 1 {
			return opt.Format.LastMinute
		}
		if minutes > 0 && minutes < 60 {
			return fmt.Sprintf(opt.Format.MinutesAgo, positive(minutes))
		}
		return Of(t, Options{
			Precision: HourPrecision,
		})
	case HourPrecision:
		hours := time.Since(t).Round(time.Hour).Hours()
		if hours > -24 && hours < -1 {
			return fmt.Sprintf(opt.Format.HoursFromNow, positive(hours))
		}
		if hours == -1 {
			return opt.Format.NextHour
		}
		if hours == 0 {
			return opt.Format.ThisHour
		}
		if hours == 1 {
			return opt.Format.LastHour
		}
		if hours > 0 && hours < 24 {
			return fmt.Sprintf(opt.Format.HoursAgo, positive(hours))
		}
		return Of(t, Options{
			Precision: DayPrecision,
		})
	case DayPrecision:
		days := time.Since(t).Round(time.Hour*24).Hours() / 24
		if days > -30 && days < -1 {
			return fmt.Sprintf(opt.Format.DaysFromNow, positive(days))
		}
		if days == -1 {
			return opt.Format.Tomorrow
		}
		if days == 0 {
			return opt.Format.Today
		}
		if days == 1 {
			return opt.Format.Yesterday
		}
		if days > 0 && days < 30 {
			return fmt.Sprintf(opt.Format.DaysAgo, positive(days))
		}
		return Of(t, Options{
			Precision: MonthPrecision,
		})
	case MonthPrecision:
		months := time.Since(t).Round(time.Hour*24*30).Hours() / (24 * 30)
		if months > -12 && months < -1 {
			return fmt.Sprintf(opt.Format.MonthsFromNow, positive(months))
		}
		if months == -1 {
			return opt.Format.NextMonth
		}
		if months == 0 {
			return opt.Format.ThisMonth
		}
		if months == 1 {
			return opt.Format.LastMonth
		}
		if months > 0 && months < 12 {
			return fmt.Sprintf(opt.Format.MonthsAgo, positive(months))
		}
		return Of(t, Options{
			Precision: YearPrecision,
		})
	default:
		// its year precision
		years := time.Since(t).Round(time.Hour*24*365).Hours() / (24 * 365)
		if years < -1 {
			return fmt.Sprintf(opt.Format.YearsFromNow, positive(years))
		}
		if years == -1 {
			return opt.Format.NextYear
		}
		if years == 0 {
			return opt.Format.ThisYear
		}
		if years == 1 {
			return opt.Format.LastYear
		}
		return fmt.Sprintf(opt.Format.YearsAgo, positive(years))
	}
}

// Format is the format of the string returned by the library.
type Format struct {
	SecondsFromNow string
	NextSecond     string
	ThisSecond     string
	SecondsAgo     string
	LastSecond     string

	MinutesFromNow string
	NextMinute     string
	ThisMinute     string
	MinutesAgo     string
	LastMinute     string

	HoursFromNow string
	NextHour     string
	ThisHour     string
	HoursAgo     string
	LastHour     string

	DaysFromNow string
	Tomorrow    string
	Today       string
	Yesterday   string
	DaysAgo     string

	MonthsFromNow string
	NextMonth     string
	ThisMonth     string
	LastMonth     string
	MonthsAgo     string

	YearsFromNow string
	NextYear     string
	ThisYear     string
	LastYear     string
	YearsAgo     string
}

// DefaultFormat is the default format of the string returned by the library.
var DefaultFormat = Format{
	SecondsFromNow: "%d seconds from now",
	NextSecond:     "1 second from now",
	ThisSecond:     "now",
	LastSecond:     "1 second ago",
	SecondsAgo:     "%d seconds ago",

	MinutesFromNow: "%d minutes from now",
	NextMinute:     "1 minute from now",
	ThisMinute:     "now",
	LastMinute:     "1 minute ago",
	MinutesAgo:     "%d minutes ago",

	HoursFromNow: "%d hours from now",
	NextHour:     "next hour",
	ThisHour:     "this hour",
	LastHour:     "last hour",
	HoursAgo:     "%d hours ago",

	DaysFromNow: "%d days from now",
	Tomorrow:    "tomorrow",
	Today:       "today",
	Yesterday:   "yesterday",
	DaysAgo:     "%d days ago",

	MonthsFromNow: "%d months from now",
	NextMonth:     "next month",
	ThisMonth:     "this month",
	LastMonth:     "last month",
	MonthsAgo:     "%d months ago",

	YearsFromNow: "%d years from now",
	NextYear:     "next year",
	ThisYear:     "this year",
	LastYear:     "last year",
	YearsAgo:     "%d years ago",
}

func positive(i float64) int {
	if i > 0 {
		return int(i)
	}
	return int(i) * -1
}
