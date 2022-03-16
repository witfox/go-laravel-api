package auth

import (
	"errors"
	"gohub/app/models/user"
	"gohub/pkg/logger"

	"github.com/gin-gonic/gin"
)

//支持 phone/email/username登录
func Attempt(account string, password string) (user.User, error) {
	_user := user.GetByMulti(account)

	if _user.ID == 0 {
		return user.User{}, errors.New("账户不存在")
	}

	if !_user.ComparePassword(password) {
		return user.User{}, errors.New("密码错误")
	}

	return _user, nil
}

func LoginByPhone(phone string) (user.User, error) {
	_user := user.GetByPhone(phone)
	if _user.ID == 0 {
		return user.User{}, errors.New("手机号未注册")
	}
	return _user, nil
}

func CurrentUser(c *gin.Context) user.User {
	userModel, ok := c.MustGet("current_user").(user.User)
	if !ok {
		logger.LogIf(errors.New("无法获取当前用户"))
		return user.User{}
	}

	return userModel
}

func CurrentUserId(c *gin.Context) string {
	return c.GetString("current_user_id")
}
