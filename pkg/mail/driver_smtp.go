package mail

import (
	"fmt"
	"gohub/pkg/logger"
	"net/smtp"

	emailPkg "github.com/jordan-wright/email"
)

type SMTP struct{}

func (s *SMTP) Send(email Email, config map[string]string) bool {
	e := emailPkg.NewEmail()
	e.From = fmt.Sprintf("%v <%v>", email.From.Name, email.From.Address)
	e.To = email.To
	e.Bcc = email.Bcc
	e.Cc = email.Cc
	e.Subject = email.Subject
	e.HTML = email.HTML

	logger.DebugJSON("发送邮件", "发送详情", e)

	err := e.Send(fmt.Sprintf("%v:%v", config["host"], config["port"]),
		smtp.PlainAuth("", config["username"], config["password"], config["host"]))
	if err != nil {
		logger.ErrorString("发送邮件", "发送邮件失败", err.Error())
		return false
	}
	logger.DebugString("发送邮件", "发送成功", "")
	return true

}
