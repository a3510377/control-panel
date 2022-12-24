package container

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CheckFromRequest check token from request
// if token is valid, set `user` info to gin.Context
// get user info from gin.Context: c.MustGet("user").(*database.DBAccount)
func (a accountContainer) CheckFromRequest(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(401, gin.H{"error": "token error"})
		c.Abort()
		return
	}

	info, status := a.db.GetUserByJWT(token)
	if status == http.StatusOK {
		c.Set("user", info)

		c.Next()
		return
	}

	// 401
	c.JSON(status, gin.H{"error": "token error"})
	c.Abort()
}
