package system

import "time"

// auto add system info cache
// return stop system info cache
func StartCacheSystemInfo(interval, max time.Duration) func() {
	maxLen := int(max / interval)
	ticker := time.NewTicker(interval)

	call := func() {
		data := SystemInfoCache{}

		end := len(SystemTimeInfo)
		if maxLen > end {
			end++
		}

		SystemTimeInfo = append([]SystemInfoCache{data}, SystemTimeInfo...)[:end]
	}

	call()
	go func() {
		for range ticker.C {
			call()
		}
	}()

	return ticker.Stop
}
