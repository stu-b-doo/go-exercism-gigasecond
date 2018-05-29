// solution to http://exercism.io/exercises/go/gigasecond/readme
package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond returns a the supplied time with a gigasecond added
func AddGigasecond(t time.Time) time.Time {

	// Convert the offset to a duration
	offset, _ := time.ParseDuration("1000000000s")

	// add the offset to the time
	return t.Add(offset)
}
