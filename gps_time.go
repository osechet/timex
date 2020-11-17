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

// ToUTC returns t as a time.Time in UTC.
func (t GpsTime) ToUTC() time.Time {
	return toUtcTime(t.Gps())
}

// Add returns the time t+d.
func (t GpsTime) Add(d time.Duration) GpsTime {
	return Gps(t.Gps() + d)
}

// Sub returns the duration t-u.
func (t GpsTime) Sub(u GpsTime) time.Duration {
	return t.Gps() - u.Gps()
}

// Equal reports whether t and u represent the same time instant.
func (t GpsTime) Equal(u GpsTime) bool {
	return t.Gps() == u.Gps()
}
