package timex

import (
	"time"
)

// The GPS datum on January 6, 1980 0:00:00. Unix time: 315964800
var gpsDatum = time.Date(1980, time.January, 6, 0, 0, 0, 0, time.UTC)

/* The year of first leap second. */
const leapSecondOrigin = 1980

var leapSequence = [][]int{
	{ /* 1980 */ 0, 0}, { /* 1981 */ 1, 1},
	{ /* 1982 */ 2, 2}, { /* 1983 */ 3, 3}, { /* 1984 */ 3, 3}, { /* 1985 */ 4, 4},
	{ /* 1986 */ 4, 4}, { /* 1987 */ 4, 5}, { /* 1988 */ 5, 5}, { /* 1989 */ 5, 6},
	{ /* 1990 */ 6, 7}, { /* 1991 */ 7, 7}, { /* 1992 */ 8, 8}, { /* 1993 */ 9, 9},
	{ /* 1994 */ 10, 10}, { /* 1995 */ 10, 11}, { /* 1996 */ 11, 11}, { /* 1997 */ 12, 12},
	{ /* 1998 */ 12, 13}, { /* 1999 */ 13, 13}, { /* 2000 */ 13, 13}, { /* 2001 */ 13, 13},
	{ /* 2002 */ 13, 13}, { /* 2003 */ 13, 13}, { /* 2004 */ 13, 13}, { /* 2005 */ 13, 14},
	{ /* 2006 */ 14, 14}, { /* 2007 */ 14, 14}, { /* 2008 */ 14, 15}, { /* 2009 */ 15, 15},
	{ /* 2010 */ 15, 15}, { /* 2011 */ 15, 15}, { /* 2012 */ 16, 16}, { /* 2013 */ 16, 16},
	{ /* 2014 */ 16, 16}, { /* 2015 */ 17, 17}, { /* 2016 */ 17, 18},
}

// leaps stores the times when leap seconds are added as Duration since the GPS origin.
var leaps []time.Duration

// Called when package is imported.
func init() {
	leaps = make([]time.Duration, 0)
	previous := 0
	for i, yearLeaps := range leapSequence {
		year := leapSecondOrigin + i
		for j, leap := range yearLeaps {
			var month time.Month
			var day int
			if j == 0 {
				month = time.June
				day = 30
			} else {
				month = time.December
				day = 31
			}
			if leap != previous {
				utcTime := time.Date(year, month, day, 23, 59, 59, 0, time.UTC)
				gpsTime := utcTime.Sub(gpsDatum) + time.Duration(leap)*time.Second
				leaps = append(leaps, gpsTime)
				previous = leap
			}
		}
	}
}

// countleaps counts the number of leaps before a given GPS time. Set toGps to true,
// if the given duration contains the leap seconds; use false otherwise.
func countleaps(gpsOffset time.Duration, toGps bool) int {
	nleaps := 0
	for i, v := range leaps {
		leap := time.Duration(v)
		// offset represents the number of leap seconds up to this date
		var offset time.Duration
		if toGps {
			offset = time.Duration(i) * time.Second
		} else {
			offset = 0
		}
		if gpsOffset >= leap-offset {
			nleaps++
		}
	}
	return nleaps
}

// ToGpsTime returns the GPS representation of the given time as a Duration from GPS origin.
func toGpsTime(t time.Time) time.Duration {
	isLeap := 0
	fromOrigin := t.Sub(gpsDatum)
	nleaps := countleaps(fromOrigin, true)
	gpsTime := fromOrigin + time.Duration(nleaps+isLeap)*time.Second
	return time.Duration(gpsTime)
}

// ToUtcTime converts the given GPS time to a UTC time.
func toUtcTime(gpsTime time.Duration) time.Time {
	t := gpsDatum.Add(time.Duration(gpsTime))
	nleaps := countleaps(gpsTime, false)
	t = t.Add(-time.Duration(nleaps) * time.Second)
	return t.UTC()
}
