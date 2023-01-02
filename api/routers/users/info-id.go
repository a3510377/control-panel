package users

import (
	"net/http"

	"github.com/a3510377/control-panel/container"
	"github.com/a3510377/control-panel/database"
	"github.com/a3510377/control-panel/service/id"
	"github.com/gin-gonic/gin"
)

const queryUserKey = "query-user"

// `/:id`
func infoByID(container *container.Container, app *gin.RouterGroup) {
	dynamicID := app.Group("/:id", func(c *gin.Context) {
		id := id.StringToID(c.Param("id"))
		if id == -1 {
			c.JSON(http.StatusNotFound, gin.H{"error": "not fond"})
			c.Abort()
			return
		}

		data := container.GetUserByID(id)
		c.Set(queryUserKey, data)
		c.Next()
	})

	dynamicID.GET("", func(c *gin.Context) {
		data := c.MustGet(queryUserKey).(*database.DBAccount)
		if data == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "not fond"})
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": data.JSON()})
	})
}
