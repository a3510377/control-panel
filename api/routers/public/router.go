package public

import (
	"github.com/a3510377/control-panel/container"
	"github.com/a3510377/control-panel/routers/public/account"
	"github.com/gin-gonic/gin"
)

func PublicRouter(container *container.Container, app *gin.RouterGroup) {
	account.AddHandler(container, app.Group("/account"))
}
