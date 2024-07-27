// Package gigasecond provides function to determine the date and time one gigasecond after a certain date
package gigasecond

// import path for the time and math packages from the standard library
import (
	"math"
	"time"
)

// AddGigasecond adds 10 ** 9 seconds to given time
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * time.Duration(math.Pow10(9)))
}
