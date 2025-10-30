package computingtime_test

import (
	"time"
)

var (
	startTime   time.Time
	elapsedTime time.Duration
)

func Start() {
	startTime = time.Now()
}

func End() time.Duration {
	elapsedTime = time.Since(startTime)
	return elapsedTime
}
