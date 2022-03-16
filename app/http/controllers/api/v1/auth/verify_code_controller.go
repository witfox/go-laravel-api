package auth

import (
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/app/requests"
	"gohub/pkg/captcha"
	"gohub/pkg/logger"
	"gohub/pkg/response"
	"gohub/pkg/verifycode"

	"github.com/gin-gonic/gin"
)

type VerifyCodeController struct {
	v1.BaseController
}

func (vc *VerifyCodeController) ShowCaptcha(c *gin.Context) {

	id, b64s, err := captcha.NewCaptcha().GenerateCaptcha()
	logger.LogIf(err)

	response.JSON(c, gin.H{
		"captcha_id":    id,
		"captcha_image": b64s,
	})
}

func (vc *VerifyCodeController) SendCodeByPhone(c *gin.Context) {

	//验证表单
	request := requests.VerifyCodeRequest{}
	if ok := requests.Validate(c, &request, requests.VerifyCodePhone); !ok {
		return
	}

	//发送短信
	if ok := verifycode.NewVerifyCode().SendSMS(request.Phone); !ok {
		response.Abort500(c, "发送短信失败")
	} else {
		response.Success(c)
	}
}

func (vc *VerifyCodeController) SendCodeByEmail(c *gin.Context) {

	request := requests.VerifyCodeEmailRequest{}
	if ok := requests.Validate(c, &request, requests.VerifyCodeEmail); !ok {
		return
	}

	if err := verifycode.NewVerifyCode().SendEmail(request.Email); err != nil {
		response.Abort500(c, "发送邮件失败")
	} else {
		response.Success(c)
	}
}
