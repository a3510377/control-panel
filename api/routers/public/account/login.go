package account

import (
	"net/http"

	"github.com/a3510377/control-panel/container"
	"github.com/a3510377/control-panel/database"
	"github.com/gin-gonic/gin"
)

func addLoginHandlers(container *container.Container, app *gin.RouterGroup) {
	app.POST("/login", func(c *gin.Context) {
		var user database.NewAccountData

		c.ShouldBindJSON(&user)
		if err := database.CheckJSONData(user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		data := container.GetUserByName(user.Username)
		if data == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "username not found"})
			return
		}
		if !data.CheckPassword(user.Password) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "password error"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": data.JSON()})
	})
}
