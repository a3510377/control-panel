package account

import (
	"net/http"

	"github.com/a3510377/control-panel/container"
	"github.com/a3510377/control-panel/database"
	"github.com/gin-gonic/gin"
)

func addCreateHandlers(container *container.Container, app *gin.RouterGroup) {
	app.POST("/create", func(c *gin.Context) {
		var newUser database.NewAccountData

		c.ShouldBindJSON(&newUser)
		data, err := container.CreateNewUser(newUser)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": "", "data": map[string]any{
			"id":         data.ID,
			"name":       data.Name,
			"permission": data.Permission,
			"created_at": data.CreatedAt,
		}})
	})
}
