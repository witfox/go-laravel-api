package v1

import (
	"gohub/pkg/auth"
	"gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

type UsersController struct {
	BaseController
}

func (uc *UsersController) CurrentUser(c *gin.Context) {
	userModel := auth.CurrentUser(c)

	response.Data(c, userModel)
}
