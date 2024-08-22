package mail

import (
	"fmt"
	mail "github.com/xhit/go-simple-mail/v2"
	"go.uber.org/zap"
	"mail-service/global"
	"mail-service/pkg/setting"
	"time"
)

type Mail struct {
	Domain     string
	Host       string
	Port       int
	Username   string
	Password   string
	Encryption string
}

type Message struct {
	From     string
	To       string
	Subject  string
	Body     string
	HTMLBody string
}

func NewMailer(config setting.MailhogSetting) *mail.SMTPClient {
	configsup := global.Config
	fmt.Println(configsup)
	server := mail.NewSMTPClient()
	server.Host = config.MailHost
	server.Port = config.MailPort
	server.Username = config.MailUsername
	server.Password = config.MailPassword
	server.Encryption = getEncryption("none")
	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second

	smtpClient, err := server.Connect()
	if err != nil {
		global.Logger.Error("smtpClient connect error", zap.Error(err))
		panic(err)
	}

	return smtpClient
}

func SendSMTPMessage(msg Message) error {
	smtpClient := global.Mailer

	email := mail.NewMSG()
	email.SetFrom(msg.From).AddTo(msg.To).SetSubject(msg.Subject)
	email.SetBody(mail.TextPlain, msg.Body)
	email.AddAlternative(mail.TextHTML, msg.HTMLBody)

	err := email.Send(smtpClient)
	if err != nil {
		global.Logger.Error("send email error", zap.Error(err))
		return err
	}

	return nil
}

func getEncryption(s string) mail.Encryption {
	switch s {
	case "tls":
		return mail.EncryptionSTARTTLS
	case "ssl":
		return mail.EncryptionSSLTLS
	case "none", "":
		return mail.EncryptionNone
	default:
		return mail.EncryptionSTARTTLS
	}
}
