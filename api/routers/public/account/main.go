package account

import (
	"fmt"

	"github.com/a3510377/control-panel/container"
	"github.com/gin-gonic/gin"
)

func checkNoToken(c *gin.Context) {
	fmt.Println("Authorization:", c.GetHeader("Authorization"))
}

func AddHandler(container *container.Container, app *gin.RouterGroup) {
	No := app.Group("/", func(ctx *gin.Context) {
		checkNoToken(ctx)
	})
	addCreateHandlers(container, No)
	addLoginHandlers(container, app)
}
