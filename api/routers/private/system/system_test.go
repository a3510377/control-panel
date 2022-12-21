package system

import (
	"testing"
	"time"
)

func TestStartCacheSystemInfo(t *testing.T) {
	stop := startCacheSystemInfo(time.Second/2, time.Second*1)
	defer stop()

	loop := true
	old := 0
	time.AfterFunc(time.Second*5, func() { loop = false })

	for loop {
		if c := len(systemTimeInfo); c != old {
			old = c

			t.Log(systemTimeInfo)
		}
	}
}
