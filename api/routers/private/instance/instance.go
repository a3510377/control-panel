package instance

import (
	"net/http"

	"github.com/a3510377/control-panel/container"
	"github.com/a3510377/control-panel/database"
	"github.com/gin-gonic/gin"
)

const instanceDataKey = "instance_data"

func addInstanceHandlers(container *container.Container, app *gin.RouterGroup) {
	app.GET("", func(c *gin.Context) {
		data := c.MustGet(instanceDataKey).(*database.DBInstance)
		c.JSON(http.StatusOK, data.Instance)
	})

	app.PATCH("", func(c *gin.Context) {
		var newInstance map[string]any

		data := c.MustGet(instanceDataKey).(*database.DBInstance)

		c.BindJSON(&newInstance)
		data.Updates(newInstance)
	})

	app.DELETE("", func(c *gin.Context) {
		data := c.MustGet(instanceDataKey).(*database.DBInstance)
		if err := data.Delete(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "instance deleted error"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"error": "instance deleted"})
	})
}
