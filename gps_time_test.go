package timex

import (
	"reflect"
	"testing"
	"time"
)

func TestGps(t *testing.T) {
	type args struct {
		offset time.Duration
	}
	tests := []struct {
		name string
		args args
		want GpsTime
	}{
		{"Epoch", args{0}, GpsTime(gpsDatum)},
		{"Inside leap table", args{948731799 * time.Second}, GpsTime(time.Date(2010, time.January, 28, 16, 36, 24, 0, time.UTC))},
		{"Before leap table", args{-315187200 * time.Second}, GpsTime(time.Date(1970, time.January, 10, 0, 0, 0, 0, time.UTC))},
		{"After leap table", args{1436486418 * time.Second}, GpsTime(time.Date(2025, time.July, 14, 0, 0, 0, 0, time.UTC))},
		{"Before leap", args{1025136014 * time.Second}, GpsTime(time.Date(2012, time.June, 30, 23, 59, 59, 0, time.UTC))},
		{"After leap", args{1025136016 * time.Second}, GpsTime(time.Date(2012, time.July, 1, 0, 0, 0, 0, time.UTC))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Gps(tt.args.offset); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Gps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGpsTime_Gps(t *testing.T) {
	tests := []struct {
		name string
		t    GpsTime
		want time.Duration
	}{
		{"GPS Datum", GpsTime(gpsDatum), 0},
		{"Inside leap table", GpsTime(time.Date(2010, time.January, 28, 16, 36, 24, 0, time.UTC)), 948731799 * time.Second},
		{"Before leap table", GpsTime(time.Date(1970, time.January, 10, 0, 0, 0, 0, time.UTC)), -315187200 * time.Second},
		{"After leap table", GpsTime(time.Date(2025, time.July, 14, 0, 0, 0, 0, time.UTC)), 1436486418 * time.Second},
		{"Before leap", GpsTime(time.Date(2012, time.June, 30, 23, 59, 59, 0, time.UTC)), 1025136014 * time.Second},
		{"After leap", GpsTime(time.Date(2012, time.July, 1, 0, 0, 0, 0, time.UTC)), 1025136016 * time.Second},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.Gps(); got != tt.want {
				t.Errorf("GpsTime.Gps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString(t *testing.T) {
	value := time.Now()
	gps := GpsTime(value)

	got := gps.String()
	want := value.String()
	if got != want {
		t.Errorf("String() = %v, want %v", got, want)
	}
}
