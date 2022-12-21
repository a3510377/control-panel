package system

import (
	"net/http"
	"time"

	"github.com/a3510377/control-panel/container"
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v3/mem"
)

type systemInfoCache struct{}

const (
	max      = time.Minute * 10 // cache 10 minutes
	interval = time.Second * 10 // interval
)

var systemTimeInfo = []systemInfoCache{}

// return stop system info cache
func AddHandler(container *container.Container, app *gin.RouterGroup) func() {
	app.GET("/mem-info", func(c *gin.Context) {
		v, err := mem.VirtualMemory()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "getting virtual memory failed"})
			return
		}

		c.JSON(http.StatusOK, v)
	})

	return startCacheSystemInfo(interval, max)
}

// auto add system info cache
// return stop system info cache
func startCacheSystemInfo(interval, max time.Duration) func() {
	maxLen := int(max / interval)
	ticker := time.NewTicker(interval)

	call := func() {
		data := systemInfoCache{}

		end := len(systemTimeInfo)
		if maxLen > end {
			end++
		}

		systemTimeInfo = append([]systemInfoCache{data}, systemTimeInfo...)[:end]
	}

	call()
	go func() {
		for range ticker.C {
			call()
		}
	}()

	return ticker.Stop
}
