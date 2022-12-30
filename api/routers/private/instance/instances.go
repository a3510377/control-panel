package instance

import (
	"net/http"

	"github.com/a3510377/control-panel/container"
	"github.com/a3510377/control-panel/database"
	"github.com/a3510377/control-panel/models"
	"github.com/gin-gonic/gin"
)

func addInstancesHandlers(container *container.Container, app *gin.RouterGroup) {
	app.POST("", func(c *gin.Context) {
		var instances models.Instance

		c.ShouldBindJSON(&instances)
		if err := database.CheckJSONData(instances); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		data, err := container.CreateNewInstance(&instances)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "create instance error"})
			return
		}
		c.JSON(200, gin.H{"data": data.JSON()})
	})
}
