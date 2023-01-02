package users

import (
	"net/http"

	"github.com/a3510377/control-panel/container"
	"github.com/a3510377/control-panel/database"
	"github.com/gin-gonic/gin"
)

// `/`
func userInfo(container *container.Container, app *gin.RouterGroup) {
	app.GET("", func(c *gin.Context) {
		users, err := container.GetUsers()
		if err != nil {
		}
		usersData := []map[string]any{}
		for _, user := range users {
			usersData = append(usersData, user.JSON())
		}
		c.JSON(http.StatusOK, gin.H{"data": usersData})
	})

	app.GET("/@me", func(c *gin.Context) {
		data := c.MustGet("user").(*database.DBAccount)

		c.JSON(http.StatusOK, gin.H{"data": data.JSON()})
	})
}
