package mail

import (
	mail "github.com/xhit/go-simple-mail/v2"
	"log"
	"mail-service/config"
	"time"
)

type Mailer struct {
	Config *config.Config
}

type Message struct {
	From     string
	To       string
	Subject  string
	Body     string
	HTMLBody string
}

func (m *Mailer) SendSMTPMessage(msg Message) error {

	server := mail.NewSMTPClient()
	server.Host = m.Config.Mailer.Host
	server.Port = m.Config.Mailer.Port
	server.Username = m.Config.Mailer.Username
	server.Password = m.Config.Mailer.Password
	server.Encryption = getEncryption(m.Config.Mailer.Encryption)
	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second

	smtpClient, err := server.Connect()
	if err != nil {
		log.Println(err)
		return err
	}

	email := mail.NewMSG()
	email.SetFrom(msg.From).AddTo(msg.To).SetSubject(msg.Subject)
	email.SetBody(mail.TextPlain, msg.Body)
	email.AddAlternative(mail.TextHTML, msg.HTMLBody)

	err = email.Send(smtpClient)
	if err != nil {
		log.Println(err)
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
