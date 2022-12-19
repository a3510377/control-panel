package container

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CheckFromRequest check token from request
// if token is valid, set `user` info to gin.Context
func (a accountContainer) CheckFromRequest(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(400, gin.H{"error": "no token"})
		c.Abort()
		return
	}

	info, status := a.db.GetUserByJWT(token)
	if status == http.StatusOK {
		c.AbortWithStatus(status)
		c.Set("user", info)

		c.Next()
		return
	}

	c.JSON(status, gin.H{"error": "token error"})
	c.Abort()
}