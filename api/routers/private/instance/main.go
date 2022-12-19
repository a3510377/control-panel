package instance

import (
	"net/http"

	"github.com/a3510377/control-panel/container"
	IDD "github.com/a3510377/control-panel/service/id"
	"github.com/gin-gonic/gin"
)

func AddHandler(container *container.Container, app *gin.RouterGroup) {
	instancePath := app.Group("/:instance_id", func(c *gin.Context) {
		id := IDD.StringToID(c.Param("instance_id"))
		// TODO add has promotion

		if id == -1 {
			c.JSON(http.StatusBadRequest, gin.H{"message": "instance_id error"})
		} else if data := container.GetInstanceByID(id); data == nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "instance_id not found"})
		} else {
			c.Set(instanceDataKey, data)
			c.Next()
			return
		}

		c.Abort()
	})
	addInstanceHandlers(container, instancePath)
	addInstancesHandlers(container, app)
}
