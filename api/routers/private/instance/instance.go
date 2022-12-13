package instance

import (
	"net/http"

	"github.com/a3510377/control-panel/database"
	bID "github.com/a3510377/control-panel/service/id"
	"github.com/gin-gonic/gin"
)

const instanceDataKey = "instance_data"

func AddHandler(db *database.DB, app *gin.RouterGroup) {
	instancePath := app.Group(":instance_id", func(c *gin.Context) {
		id := bID.StringToID(c.Param("instance_id"))
		// TODO add has promotion

		if id == -1 {
			c.JSON(http.StatusBadRequest, gin.H{"message": "instance_id error"})
		} else if data := db.GetInstanceByID(id); data == nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "instance_id not found"})
		} else {
			c.Set(instanceDataKey, data)
			return
		}
		c.Abort()
	})

	instancePath.GET("/", func(c *gin.Context) {
		data := c.MustGet(instanceDataKey).(*database.DBInstance)
		c.JSON(http.StatusOK, data.Instance)
	})

	instancePath.POST("/updata", func(c *gin.Context) {
	})
}
