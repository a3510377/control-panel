package routers

import (
	"github.com/a3510377/control-panel/container"
	"github.com/a3510377/control-panel/routers/account"
	"github.com/a3510377/control-panel/routers/overview"
	"github.com/a3510377/control-panel/routers/users"
	"github.com/gin-gonic/gin"
)

func loadRoutes(container *container.Container, app *gin.RouterGroup) {
	account.AddHandler(container, app.Group("/account"))
	overview.AddHandler(container, app.Group("/overview", container.Account.CheckFromRequest))
	users.AddHandler(container, app.Group("/users", container.Account.CheckFromRequest))
}
