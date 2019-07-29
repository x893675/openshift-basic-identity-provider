package mail

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/smtp"
	"openshift-basic-identity-provider/helper"
)

var emailcfg *emailConfig

type emailConfig struct {
	host      string
	port      int
	email     string
	password  string
	emailFrom string
}

func init() {
	emailcfg = &emailConfig{}
	helper.SetLocalVar("MAIL_HOST", &emailcfg.host, "")
	var port string
	helper.SetLocalVar("MAIL_PORT", &port, "")
	emailcfg.port, _ = helper.S(port).Int()
	helper.SetLocalVar("MAIL_ADDR", &emailcfg.email, "")
	helper.SetLocalVar("MAIL_PASSWORD", &emailcfg.password, "")
	helper.SetLocalVar("MAIL_MAILFORM", &emailcfg.emailFrom, "")
	fmt.Println(emailcfg)
}

// SendMail 发送邮件
func SendMail(toEmail, subject, content string) error {
	headers := make(map[string]string)
	headers["From"] = emailcfg.emailFrom + "<" + emailcfg.email + ">"
	headers["To"] = toEmail
	headers["Subject"] = subject
	headers["Content-Type"] = "text/html; charset=UTF-8"

	message := ""
	for key, value := range headers {
		message += fmt.Sprintf("%s: %s\r\n", key, value)
	}
	message += "\r\n" + content

	auth := smtp.PlainAuth("", emailcfg.email, emailcfg.password, emailcfg.host)

	err := sendMailUsingTLS(
		fmt.Sprintf("%s:%d", emailcfg.host, emailcfg.port),
		auth,
		emailcfg.email,
		[]string{toEmail},
		message,
	)
	return err
}

//参考net/smtp的func SendMail()
//使用net.Dial连接tls(ssl)端口时, smtp.NewClient()会卡住且不提示err
//len(to) > 1 时, to[1]开始提示是密送
func sendMailUsingTLS(addr string, auth smtp.Auth, from string,
	to []string, message string) error {

	client, err := createSMTPClient(addr)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer client.Close()

	if auth != nil {
		if ok, _ := client.Extension("AUTH"); ok {
			if err := client.Auth(auth); err != nil {
				log.Println(err.Error())
				return err
			}
		}
	}

	if err := client.Mail(from); err != nil {
		return err
	}

	for _, addr := range to {
		if err := client.Rcpt(addr); err != nil {
			return err
		}
	}

	writeCloser, err := client.Data()
	if err != nil {
		return err
	}

	_, err = writeCloser.Write([]byte(message))
	if err != nil {
		return err
	}

	err = writeCloser.Close()

	if err != nil {
		return err
	}

	return client.Quit()
}

func createSMTPClient(addr string) (*smtp.Client, error) {
	conn, err := tls.Dial("tcp", addr, nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	host, _, _ := net.SplitHostPort(addr)
	return smtp.NewClient(conn, host)
}
