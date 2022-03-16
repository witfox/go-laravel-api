package sms

import (
	"gohub/pkg/config"
	"sync"
)

type Message struct {
	Template string
	Data     map[string]string

	Content string
}

//发送短信操作类
type SMS struct {
	Driver Driver
}

//单例模式
var once sync.Once

var intervalSMS *SMS

func NewSMS() *SMS {
	once.Do(func() {
		intervalSMS = &SMS{
			Driver: &Aliyun{},
		}
	})
	return intervalSMS
}

func (sms *SMS) Send(phone string, message Message) bool {
	return sms.Driver.Send(phone, message, config.GetStringMapString("sms.aliyun"))
}
