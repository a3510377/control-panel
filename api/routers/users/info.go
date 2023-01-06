package users

import (
	"net/http"
	"strconv"

	"github.com/a3510377/control-panel/container"
	"github.com/a3510377/control-panel/database"
	"github.com/gin-gonic/gin"
)

// `/`
func userInfo(container *container.Container, app *gin.RouterGroup) {
	app.GET("", func(c *gin.Context) {
		// TODo
		page, _ := strconv.Atoi(c.Query("page"))
		pageSize, _ := strconv.Atoi(c.Query("page_size"))
		if pageSize > 100 {
			// page size limit 100
			pageSize = 100
		} else if pageSize <= 0 {
			pageSize = 10
		}

		container.DB.Limit(pageSize).Offset((page - 1) * pageSize).Model(&database.DBAccount{}).Find(&database.DBAccount{})

		users, err := container.GetUsers()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "get users failed"})
			return
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
