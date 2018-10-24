// stat.go.

// Application's Statistics.

package stat

import (
	"time"
)

const TimeFormat = "2006-01-02 15:04:05 MST"

var StartTime time.Time
var StartTimestamp int64

var StopTime time.Time
var StopTimestamp int64

// Initializes Statistics.
func Init() error {

	StartTime = time.Now()
	StartTimestamp = StartTime.Unix()

	return nil
}

// Finalizes Statistics.
func Fin() error {

	StopTime = time.Now()
	StopTimestamp = StopTime.Unix()

	return nil
}

// Returns the Duration (in Seconds) of the Service being alive ("up-time").
func GetTimeBeingAlive() int64 {

	var tsNow int64
	var upTime int64

	tsNow = time.Now().Unix()
	upTime = tsNow - StartTimestamp

	return upTime
}
