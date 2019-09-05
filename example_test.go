package timex_test

import (
	"fmt"
	"time"

	"github.com/osechet/timex"
)

// Display the GPS time of the given time, in microseconds.
func ExampleGpsTime() {
	fmt.Println(int64(timex.GpsTime(time.Date(2010, time.January, 28, 16, 36, 24, 0, time.UTC)).Gps() / time.Microsecond))
	// Output: 948731799000000
}
