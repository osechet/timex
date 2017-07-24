package timex

import (
	"reflect"
	"testing"
	"time"
)

func Test_init(t *testing.T) {
	expectedLeaps := []time.Duration{
		/* Jun 1981 */ 46828800000 * time.Millisecond,
		/* Jun 1982 */ 78364801000 * time.Millisecond,
		/* Jun 1983 */ 109900802000 * time.Millisecond,
		/* Jun 1985 */ 173059203000 * time.Millisecond,
		/* Dec 1987 */ 252028804000 * time.Millisecond,
		/* Dec 1989 */ 315187205000 * time.Millisecond,
		/* Dec 1990 */ 346723206000 * time.Millisecond,
		/* Jun 1992 */ 393984007000 * time.Millisecond,
		/* Jun 1993 */ 425520008000 * time.Millisecond,
		/* Jun 1994 */ 457056009000 * time.Millisecond,
		/* Dec 1995 */ 504489610000 * time.Millisecond,
		/* Jun 1997 */ 551750411000 * time.Millisecond,
		/* Dec 1998 */ 599184012000 * time.Millisecond,
		/* Dec 2005 */ 820108813000 * time.Millisecond,
		/* Dec 2008 */ 914803214000 * time.Millisecond,
		/* Jun 2012 */ 1025136015000 * time.Millisecond,
		/* Jun 2015 */ 1119744016000 * time.Millisecond,
		/* Dec 2016 */ 1167264017000 * time.Millisecond,
	}
	if !reflect.DeepEqual(leaps, expectedLeaps) {
		t.Errorf("invalid leaps = %v, want %v", leaps, expectedLeaps)
	}
}

func Test_countleaps(t *testing.T) {
	type args struct {
		gpsTime time.Duration
		toGps   bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"GPS Datum", args{time.Duration(0), false}, 0},
		{"First Leap", args{46915200 * time.Second, false}, 1},
		{"Second Leap", args{78451201 * time.Second, false}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countleaps(tt.args.gpsTime, tt.args.toGps); got != tt.want {
				t.Errorf("countleaps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToGpsTime(t *testing.T) {
	type args struct {
		utcTime time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Duration
	}{
		{"GPS Datum", args{gpsDatum}, 0},
		{"Inside leap table", args{time.Date(2010, time.January, 28, 16, 36, 24, 0, time.UTC)}, 948731799 * time.Second},
		{"Before leap table", args{time.Date(1970, time.January, 10, 0, 0, 0, 0, time.UTC)}, -315187200 * time.Second},
		{"After leap table", args{time.Date(2025, time.July, 14, 0, 0, 0, 0, time.UTC)}, 1436486418 * time.Second},
		{"Before leap", args{time.Date(2012, time.June, 30, 23, 59, 59, 0, time.UTC)}, 1025136014 * time.Second},
		{"After leap", args{time.Date(2012, time.July, 1, 0, 0, 0, 0, time.UTC)}, 1025136016 * time.Second},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toGpsTime(tt.args.utcTime); got != tt.want {
				t.Errorf("ToGpsTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUtcTime(t *testing.T) {
	type args struct {
		gpsTime time.Duration
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{"Epoch", args{0}, gpsDatum},
		{"Inside leap table", args{948731799 * time.Second}, time.Date(2010, time.January, 28, 16, 36, 24, 0, time.UTC)},
		{"Before leap table", args{-315187200 * time.Second}, time.Date(1970, time.January, 10, 0, 0, 0, 0, time.UTC)},
		{"After leap table", args{1436486418 * time.Second}, time.Date(2025, time.July, 14, 0, 0, 0, 0, time.UTC)},
		{"Before leap", args{1025136014 * time.Second}, time.Date(2012, time.June, 30, 23, 59, 59, 0, time.UTC)},
		{"After leap", args{1025136016 * time.Second}, time.Date(2012, time.July, 1, 0, 0, 0, 0, time.UTC)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toUtcTime(tt.args.gpsTime); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToUtcTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUtcGpsUtc(t *testing.T) {
	utcTime := time.Date(2016, time.October, 12, 9, 26, 11, 0, time.UTC)
	result := toUtcTime(toGpsTime(utcTime))
	if result != utcTime {
		t.Errorf("ToUtcTime(ToGpsTime()) = %v, want %v", result, utcTime)
	}
}

func TestGpsUtcGps(t *testing.T) {
	gpsTime := 1158536014 * time.Second
	result := toGpsTime(toUtcTime(gpsTime))
	if result != gpsTime {
		t.Errorf("ToGpsTime(ToUtcTime()) = %v, want %v", result, gpsTime)
	}
}
