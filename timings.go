package timings

import (
	"log"
	"os"
	"time"
)

// Tracker is a simple device to track timings
type Tracker struct {
	start time.Time
	end   time.Time
}

// Start records the start time
func (t *Tracker) Start() {
	t.start = time.Now()
}

// Stop records the end time
func (t *Tracker) Stop() {
	t.end = time.Now()
}

// Since returns a Duration since the start time
func (t *Tracker) Since() time.Duration {
	return time.Since(t.start)
}

// Duration returns the Duration of time between Start() and Stop()
func (t *Tracker) Duration() time.Duration {
	return t.end.Sub(t.start)
}

// Track tracks how long a function takes, to a log.Logger
// e.g. make the first line
// defer Track("func_name", time.Now(), logger).
func Track(name string, start time.Time, l *log.Logger) {
	elapsed := time.Since(start)
	l.Printf("%s took %s", name, elapsed)
}

// TrackOut tracks how long a function takes, to Stdout
// e.g. make the first line:
// defer TrackOut("func_name", time.Now()).
func TrackOut(name string, start time.Time) {
	Track(name, start, Stdout)
}

// TrackErr tracks how long a function takes, to Stderr
// e.g. make the first line:
// defer Trackerr("func_name", time.Now()).
func TrackErr(name string, start time.Time) {
	Track(name, start, Stderr)
}

// TrackIfLate is the same as Track(..), but only prints if the elapsed
// time is greater than specified.
func TrackIfLate(name string, start time.Time, late time.Duration, l *log.Logger) {
	elapsed := time.Since(start)
	if elapsed > late {
		l.Printf("%s took %s", name, elapsed)
	}
}

// Now is equivalent to time.Now(), but saves the import of "time" if
// it's not otherwise needed
func Now() time.Time {
	return time.Now()
}

var (
	// Stdout is a log.Logger that outputs formatted messages to os.Stdout, for convenience.
	Stdout = log.New(os.Stdout, "[TIMING]", log.Lshortfile)
	// Stderr is a log.Logger that outputs formattedd messages to os.Stderr, for convenience.
	Stderr = log.New(os.Stderr, "[TIMING]", log.Lshortfile)
)
