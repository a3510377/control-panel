package account

import (
	"net/http"

	"github.com/a3510377/control-panel/container"
	"github.com/gin-gonic/gin"
)

type CreateAccountRequestData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func addCreateHandlers(container *container.Container, app *gin.RouterGroup) {
	app.POST("/create", func(c *gin.Context) {
		var newUser CreateAccountRequestData

		if err := c.ShouldBindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		container.CreateNewUser("test", "test")
	})
}
