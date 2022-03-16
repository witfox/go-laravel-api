package auth

import (
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/app/models/user"
	"gohub/app/requests"
	"gohub/pkg/jwt"
	"gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

type SignupController struct {
	v1.BaseController
}

func (sc *SignupController) IsPhoneExists(c *gin.Context) {

	request := requests.PhoneExistRequest{}
	if ok := requests.Validate(c, &request, requests.ValidateSignupPhone); !ok {
		return
	}

	response.JSON(c, gin.H{
		"exists": user.IsPhoneExists(request.Phone),
	})
}

func (sc *SignupController) IsEmailExists(c *gin.Context) {
	request := requests.EmailExistRequest{}

	if ok := requests.Validate(c, &request, requests.ValidateSignupEamil); !ok {
		return
	}
	response.JSON(c, gin.H{
		"exists": user.IsEmailExists(request.Email),
	})
}

func (sc *SignupController) SignupUsingPhone(c *gin.Context) {
	request := requests.SignupUsingPhoneRequest{}

	if ok := requests.Validate(c, &request, requests.SignupUsingPhone); !ok {
		return
	}

	//创建用户
	_user := user.User{
		Name:     request.Name,
		Phone:    request.Phone,
		Password: request.Password,
	}
	_user.Create()

	if _user.ID > 0 {
		token := jwt.NewJWT().IssueToken(_user.GetStringID(), _user.Name)
		response.CreatedJSON(c, gin.H{
			"data":  _user,
			"token": token,
		})
	} else {
		response.Abort500(c, "用户注册失败，请重新尝试")
	}

}

func (sc *SignupController) SignupUsingEmail(c *gin.Context) {
	request := requests.SignupUsingEmailRequest{}

	if ok := requests.Validate(c, &request, requests.SignupUsingEmail); !ok {
		return
	}

	_user := user.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
	_user.Create()

	if _user.ID > 0 {
		token := jwt.NewJWT().IssueToken(_user.GetStringID(), _user.Name)
		response.CreatedJSON(c, gin.H{
			"data":  _user,
			"token": token,
		})
	} else {
		response.Abort500(c, "创建用户失败，请稍后尝试~")
	}
}
