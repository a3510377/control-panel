package users

import (
	"net/http"

	"github.com/a3510377/control-panel/container"
	"github.com/a3510377/control-panel/database"
	"github.com/gin-gonic/gin"
)

// `/`
func userInfo(container *container.Container, app *gin.RouterGroup) {
	checkToken := app.Group("", container.Account.CheckFromRequest)

	checkToken.GET("/@me", func(c *gin.Context) {
		data := c.MustGet("user").(*database.DBAccount)

		c.JSON(http.StatusOK, gin.H{"data": data.JSON()})
	})
}
