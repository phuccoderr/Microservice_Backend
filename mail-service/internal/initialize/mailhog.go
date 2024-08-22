package initialize

import (
	"mail-service/global"
	"mail-service/internal/mail"
)

func InitMailhog() {
	mailer := mail.NewMailer(global.Config.Mailhog)
	global.Mailer = mailer
}
