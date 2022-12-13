package privateRouter

import (
	"github.com/a3510377/control-panel/database"
	"github.com/a3510377/control-panel/routers/private/instance"
	"github.com/gin-gonic/gin"
)

func PrivateRouter(db *database.DB, app *gin.RouterGroup) {
	instanceRouter := app.Group("/instance")
	instance.AddHandler(db, instanceRouter)
}
