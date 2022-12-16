package instance

import (
	"github.com/a3510377/control-panel/container"
	"github.com/gin-gonic/gin"
)

func addInstancesHandlers(container *container.Container, app *gin.RouterGroup) {
	app.POST("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "POST"})
	})
}
