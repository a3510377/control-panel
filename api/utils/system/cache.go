package system

import "time"

const (
	CacheMax      = time.Minute * 10 // cache 10 minutes
	CacheInterval = time.Second * 10 // interval
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
