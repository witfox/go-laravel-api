package validators

import (
	"gohub/pkg/captcha"
	"gohub/pkg/verifycode"
)

func ValidateCaptcha(captchaID, captchaAnswer string, errs map[string][]string) map[string][]string {
	if ok := captcha.NewCaptcha().VerifyCaptcha(captchaID, captchaAnswer); !ok {
		errs["captcha_answer"] = append(errs["captcha_answer"], "图片验证码错误")
	}
	return errs
}

func ValidatePasswordConfirm(password string, passwordConfirm string, errs map[string][]string) map[string][]string {
	if password != passwordConfirm {
		errs["password_confirm"] = append(errs["password_confirm"], "两次密码输入不一致")
	}
	return errs
}

func ValidateVerifyCode(phone string, code string, errs map[string][]string) map[string][]string {
	if ok := verifycode.NewVerifyCode().CheckAnswer(phone, code); !ok {
		errs["verify_code"] = append(errs["verify_code"], "手机验证码错误")
	}
	return errs
}
