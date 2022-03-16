package requests

import (
	"fmt"
	"gohub/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type ValidatorFunc func(interface{}, *gin.Context) map[string][]string

func validate(data interface{}, rules govalidator.MapData, message govalidator.MapData) map[string][]string {
	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		TagIdentifier: "valid",
		Messages:      message,
	}
	return govalidator.New(opts).ValidateStruct()
}

func Validate(c *gin.Context, obj interface{}, handler ValidatorFunc) bool {
	//解析json请求 支持json，表单，地址参数
	if err := c.ShouldBind(&obj); err != nil {
		response.BadRequest(c, err)

		fmt.Println(err)
		return false
	}

	//表单验证
	errs := handler(obj, c)
	if len(errs) > 0 {
		response.ValidationError(c, errs)
		return false
	}
	return true
}
