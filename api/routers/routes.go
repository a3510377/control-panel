package routers

import (
	"github.com/a3510377/control-panel/container"
	"github.com/a3510377/control-panel/routers/account"
	"github.com/a3510377/control-panel/routers/instances"
	"github.com/a3510377/control-panel/routers/overview"
	"github.com/a3510377/control-panel/routers/users"
	"github.com/gin-gonic/gin"
)

func loadRoutes(container *container.Container, app *gin.RouterGroup) {
	apiRouter := app.Group("/api")

	account.AddHandler(container, apiRouter.Group("/account"))

	overview.AddHandler(container, apiRouter.Group("/overview", container.Account.CheckFromRequest))
	users.AddHandler(container, apiRouter.Group("/users", container.Account.CheckFromRequest))
	instances.AddHandler(container, apiRouter.Group("/instances", container.Account.CheckFromRequest))
}
