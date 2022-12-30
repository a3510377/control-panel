package system

import (
	"time"
)

const (
	CacheMax      = time.Minute * 10 // cache 10 minutes
	CacheInterval = time.Second * 10 // interval
	MaxLen        = int(CacheMax / CacheInterval)
)

var SystemTimeInfo = []SystemInfo{}

// auto add system info cache
// return stop system info cache
func StartCacheSystemInfo(interval, max time.Duration) func() {
	maxLen := int(max / interval)
	ticker := time.NewTicker(interval)

	call := func() {
		data := GetNowSystemInfo()

		end := len(SystemTimeInfo)
		if maxLen > end {
			end++
		}

		SystemTimeInfo = append([]SystemInfo{data}, SystemTimeInfo...)[:end]
	}

	call()
	go func() {
		for range ticker.C {
			call()
		}
	}()

	return ticker.Stop
}

// wait for catch system info
func WaitNowSystemInfo() []SystemInfo {
	for start := time.Now(); len(SystemTimeInfo) > 0 && time.Since(start) < time.Second*3; {
		return SystemTimeInfo
	}

	return SystemTimeInfo // is no use
}
