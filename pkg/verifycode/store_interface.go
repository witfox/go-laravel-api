package verifycode

//验证码存储方式接口
type Store interface {
	Set(id string, value string) bool

	Get(id string, clear bool) string

	Verify(id, answer string, clear bool) bool
}
