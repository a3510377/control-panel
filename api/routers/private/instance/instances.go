package instance

import (
	"github.com/a3510377/control-panel/container"
	"github.com/a3510377/control-panel/models"
	"github.com/gin-gonic/gin"
)

func addInstancesHandlers(container *container.Container, app *gin.RouterGroup) {
	app.POST("/", func(c *gin.Context) {
		var instances models.Instance

		err := c.ShouldBindJSON(&instances)
		if err != nil {
			c.JSON(500, gin.H{})
		}
		container.CreateNewInstance(&models.Instance{})
		c.JSON(200, gin.H{"message": "POST"})
	})
}
