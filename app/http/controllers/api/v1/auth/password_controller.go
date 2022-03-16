package auth

import (
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/app/models/user"
	"gohub/app/requests"
	"gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

type PasswordController struct {
	v1.BaseController
}

//重置密码
func (pc *PasswordController) ResetByPhone(c *gin.Context) {

	//验证请求
	request := requests.ResetByPhoneRequest{}
	if ok := requests.Validate(c, &request, requests.ResetByPhone); !ok {
		return
	}

	//更新密码
	userModel := user.GetByPhone(request.Phone)
	if userModel.ID == 0 {
		response.Abort404(c)
	} else {
		userModel.Password = request.Password
		userModel.Save()

		response.Success(c)
	}

}

func (pc *PasswordController) ResetByEmail(c *gin.Context) {
	//验证请求
	request := requests.ResetByEmailRequest{}
	if ok := requests.Validate(c, &request, requests.ResetByEmail); !ok {
		return
	}

	//更新密码
	userModel := user.GetByEamil(request.Email)
	if userModel.ID == 0 {
		response.Abort404(c)
	} else {
		userModel.Password = request.Password
		userModel.Save()

		response.Success(c)
	}
}
