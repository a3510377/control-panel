package account

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/a3510377/control-panel/container"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CreateAccountRequestData struct {
	Username string `json:"username" validate:"required,min=5,max=20"`
	Password string `json:"password" validate:"required,min=5,max=20"`
}

func addCreateHandlers(container *container.Container, app *gin.RouterGroup) {
	validate := validator.New()

	app.POST("/create", func(c *gin.Context) {
		var newUser CreateAccountRequestData

		c.ShouldBindJSON(&newUser)
		err := validate.Struct(newUser)
		if err != nil {
			for _, err := range err.(validator.ValidationErrors) {
				errorMsg := strings.ToLower(err.Field())

				switch err.Tag() {
				case "required":
					errorMsg += " is required."
				case "min":
					errorMsg += fmt.Sprintf(" (%v) required at least %v", err.Type(), err.Param())
				case "max":
					errorMsg += fmt.Sprintf(" (%v) only has a maximum of %v", err.Type(), err.Param())
				default:
					errorMsg = err.Error()
				}

				c.JSON(http.StatusBadRequest, gin.H{"error": errorMsg})
				return
			}
		}
		container.CreateNewUser(newUser.Username, newUser.Password)
		c.JSON(http.StatusOK, gin.H{"token": ""})
	})
}
