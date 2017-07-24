package timex

import (
	"time"
)

// A GpsTime represents an instant in time with nanosecond precision.
type GpsTime time.Time

// Gps returns the local Time corresponding to the given GPS offset.
func Gps(offset time.Duration) GpsTime {
	return GpsTime(toUtcTime(offset))
}

// Gps returns t as a GPS time, a Duration from GPS epoch.
func (t GpsTime) Gps() time.Duration {
	value := toGpsTime(time.Time(t))
	return time.Duration(value)
}

func (t GpsTime) String() string {
	return time.Time(t).String()
}
