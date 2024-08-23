package mail

import (
	mail "github.com/xhit/go-simple-mail/v2"
	"go.uber.org/zap"
	"mail-service/global"
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

func SendSMTPMessage(msg Message) error {
	config := global.Config.Mailhog

	server := mail.NewSMTPClient()
	server.Host = config.MailHost
	server.Port = config.MailPort
	server.Username = config.MailUsername
	server.Password = config.MailPassword
	server.Encryption = getEncryption(config.MailEncryption)
	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second

	email := mail.NewMSG()
	email.SetFrom(msg.From).AddTo(msg.To).SetSubject(msg.Subject)
	email.SetBody(mail.TextPlain, msg.Body)
	email.AddAlternative(mail.TextHTML, msg.HTMLBody)

	smtpClient, err := server.Connect()
	if err != nil {
		global.Logger.Error("smtpClient connect error", zap.Error(err))
		panic(err)
	}

	err = email.Send(smtpClient)
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
