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

		token, status := data.CreateNewJWT()
		if status != http.StatusOK {
			c.JSON(status, gin.H{"error": "create token error"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": data.JSON(), "token": token})
	})
	app.Use(container.Account.CheckFromRequest).GET("/@me", func(c *gin.Context) {
		data := c.MustGet("user").(*database.DBAccount)

		c.JSON(http.StatusOK, gin.H{"data": data.JSON()})
	})
	app.GET("/:id") // TODO
}
