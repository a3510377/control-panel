package system

import (
	"net/http"

	"github.com/a3510377/control-panel/container"
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v3/mem"
)

// return stop system info cache
func AddHandler(container *container.Container, app *gin.RouterGroup) {
	app.GET("/mem-info", func(c *gin.Context) {
		v, err := mem.VirtualMemory()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "getting virtual memory failed"})
			return
		}

		c.JSON(http.StatusOK, v)
	})

	return
}
