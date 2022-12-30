package system

import (
	"net/http"
	"strconv"

	"github.com/a3510377/control-panel/container"
	"github.com/a3510377/control-panel/utils/system"
	"github.com/gin-gonic/gin"
)

// return stop system info cache
func AddHandler(container *container.Container, app *gin.RouterGroup) {
	system.StartCacheSystemInfo(system.CacheInterval, system.CacheMax) // return stop func

	app.GET("", func(c *gin.Context) {
		if limitStr := c.Query("limit"); limitStr != "" {
			limit, err := strconv.Atoi(limitStr)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "limit must be a number",
				})
				return
			}
			if limit <= 0 || limit > system.MaxLen {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "limit must be between 0 and " + strconv.Itoa(system.MaxLen),
				})
				return
			}

			data := system.WaitNowSystemInfo()
			if len(data) < limit {
				limit = len(data)
			}

			c.JSON(http.StatusOK, data[:limit])
			return
		}

		c.JSON(http.StatusOK, system.WaitNowSystemInfo())
	})

	app.GET("/base", func(c *gin.Context) {
		c.JSON(http.StatusOK, system.GetSystemBaseInfo())
	})
}
