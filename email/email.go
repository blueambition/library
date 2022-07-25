package email

import (
	"fmt"
	"net/smtp"
	"strings"
)

//发送邮件
//host smtp.exmail.qq.com
//sender support@trustex.club
//port 端口号 587
//password Youshiqingda110
//nickname 发送人昵称
//toEmails 收件人列表
//subject 主题
//body 内容
//contentType 类型
func Send(host string, port string, sender string, password string, nickname string, toEmails []string, subject string, body string, contentType string) error {
	//auth :=smtp.CRAMMD5Auth(sender, password)
	auth := smtp.PlainAuth("", sender, password, host)
	if contentType == "html" {
		contentType = "Content-Type: text/html; charset=UTF-8"
	} else {
		contentType = "Content-Type: text/plain; charset=UTF-8"
	}
	msg := []byte("To: " + strings.Join(toEmails, ",") + "\r\nFrom: " + nickname +
		"<" + sender + ">\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + body)
	err := smtp.SendMail(host+":"+port, auth, sender, toEmails, msg)

	if err != nil {
		fmt.Println("[Email]", string(msg)+"\r\n"+err.Error())
		return err
	}

	return nil
}
