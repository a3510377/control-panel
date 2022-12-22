package system

import (
	"net/http"
	"time"

	"github.com/a3510377/control-panel/container"
	"github.com/a3510377/control-panel/utils/system"
	"github.com/gin-gonic/gin"
)

// return stop system info cache
func AddHandler(container *container.Container, app *gin.RouterGroup) {
	system.StartCacheSystemInfo(system.CacheInterval, system.CacheMax) // return stop func

	app.GET("/", func(c *gin.Context) {
		for start := time.Now(); len(system.SystemTimeInfo) > 0 && time.Since(start) < time.Second*3; { // wait for catch
			c.JSON(http.StatusOK, system.SystemTimeInfo[0])

			break
		}
	})
}
