//短信服务商接口
package sms

type Driver interface {
	Send(phone string, message Message, config map[string]string) bool
}
