package utils

import (
	"fmt"
	"github.com/astaxie/beego/utils"
	"regexp"
	"strings"
)

// SendEmail 使用beego下utils下的NewEMail
func SendEmail(to_email, msg string, username, password, host, port string) error {
	emailConfig := fmt.Sprintf(`{"username":"%s","password":"%s","host":"%s","port":%s}`, username, password, host, port)
	emailConn := utils.NewEMail(emailConfig) // beego下的
	emailConn.From = strings.TrimSpace(username)
	emailConn.To = []string{strings.TrimSpace(to_email)}
	emailConn.Subject = "ChainLogix"
	//注意这里我们发送给用户的是激活请求地址
	emailConn.Text = msg
	err := emailConn.Send()
	if err != nil {
		return err
	}
	return nil
}

// ValidateEmail 使用正则表达式检查邮箱地址是否有效
func ValidateEmail(email string) bool {
	// 定义电子邮件的正则表达式
	var emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

	// 使用正则表达式匹配邮箱地址
	return emailRegex.MatchString(email)
}
