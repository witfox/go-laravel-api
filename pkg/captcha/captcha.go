package captcha

import (
	"gohub/pkg/app"
	"gohub/pkg/config"
	"gohub/pkg/redis"
	"sync"

	"github.com/mojocn/base64Captcha"
)

type Captcha struct {
	Base64Captcha *base64Captcha.Captcha
}

var once sync.Once

var internalCaptcha *Captcha

//单例模式获取
func NewCaptcha() *Captcha {
	once.Do(func() {
		internalCaptcha = &Captcha{}
		store := RedisStore{
			RedisClient: redis.Redis,
			KeyPrefix:   config.GetString("app.name") + ":captcha:",
		}

		driver := base64Captcha.NewDriverDigit(
			config.GetInt("captcha.height"),      // 宽
			config.GetInt("captcha.width"),       // 高
			config.GetInt("captcha.length"),      // 长度
			config.GetFloat64("captcha.maxskew"), // 数字的最大倾斜角度
			config.GetInt("captcha.dotcount"),    // 图片背景里的混淆点数量
		)

		internalCaptcha.Base64Captcha = base64Captcha.NewCaptcha(driver, &store)
	})

	return internalCaptcha
}

//生成图片验证码
func (c *Captcha) GenerateCaptcha() (id string, b64s string, err error) {
	return c.Base64Captcha.Generate()
}

//验证
func (c *Captcha) VerifyCaptcha(id, answer string) bool {

	// 方便本地和 API 自动测试
	if !app.IsProduction() && id == config.GetString("captcha.testing_key") {
		return true
	}

	return c.Base64Captcha.Verify(id, answer, false)
}
