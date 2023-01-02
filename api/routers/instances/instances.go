package instances

import (
	"github.com/a3510377/control-panel/container"
	"github.com/gin-gonic/gin"
)

func instances(container *container.Container, app *gin.RouterGroup) {
	app.GET("", func(c *gin.Context) {
		// TODO add instances
		// data := c.MustGet("user").(*database.DBAccount)
	})
}
