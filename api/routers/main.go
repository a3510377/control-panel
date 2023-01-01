package routers

import (
	"strings"
	"time"

	"github.com/a3510377/control-panel/container"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type RouterConfig struct {
	NoRouteHandlers []gin.HandlerFunc
}

func Routers(container *container.Container, config RouterConfig) *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(corsConfig()))

	router.NoRoute(append([]gin.HandlerFunc{
		func(c *gin.Context) {
			path := c.Request.URL.Path
			if strings.HasPrefix(path, "/api") {
				c.JSON(404, gin.H{"error": "not found"})
				c.Abort()
			}
		},
	}, config.NoRouteHandlers...)...)

	loadRoutes(container, router.Group(""))

	return router
}

func corsConfig() cors.Config {
	config := cors.DefaultConfig()

	if gin.Mode() == gin.DebugMode {
		config.AllowAllOrigins = true
		config.AllowMethods = []string{"GET", "POST", "DELETE", "OPTIONS", "PUT"}
		config.AllowHeaders = []string{
			"Authorization", "Content-Type", "Upgrade", "Origin",
			"Connection", "Accept-Encoding", "Accept-Language", "Host",
			"Access-Control-Request-Method", "Access-Control-Request-Headers",
		}
	} else {
		config.AllowMethods = []string{"GET", "POST", "DELETE", "OPTIONS", "PUT"}
		config.AllowHeaders = []string{
			"Authorization", "Content-Type", "Origin",
			"Connection", "Accept-Encoding", "Accept-Language", "Host",
		}
		config.AllowOrigins = []string{"http://localhost:8080"} // TODO add config.AllowOrigins
	}

	config.MaxAge = 1 * time.Hour
	config.AllowCredentials = true
	config.ExposeHeaders = []string{"Content-Length"}

	return config
}
